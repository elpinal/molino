package molino

type PersistentVector struct {
	cnt   int
	shift uint
	root  PersistentVector_Node
	tail  []interface{}
}

type PersistentVector_Node struct {
	//edit
	array []interface{}
}

type TransientVector struct {
	cnt   int
	shift uint
	root  PersistentVector_Node
	tail  []interface{}
}

type ChunkedSeq struct {
	ASeq
	vec    PersistentVector
	node   []interface{}
	i      int
	offset int
}

var PersistentVector_EMPTY_NODE = PersistentVector_Node{array: make([]interface{}, 0, 8)}

var PersistentVector_EMPTY = PersistentVector{cnt: 0, shift: 5, root: PersistentVector_EMPTY_NODE}

func (v PersistentVector) create(obj interface{}) PersistentVector {
	if items, ok := obj.(ISeq); ok {
		var ret = PersistentVector_EMPTY.asTransient()
		for ; items != nil; items = items.next() {
			ret = ret.conj(items.first())
		}
		return ret.persistent()
	} else if items, ok := obj.(List); ok {
		var ret = PersistentVector_EMPTY.asTransient()
		for _, item := range items {
			ret = ret.conj(item)
		}
		return ret.persistent()
	}
	panic("can't create PersistentVector")
}

func (v PersistentVector) asTransient() TransientVector {
	return TransientVector{cnt: v.cnt, shift: v.shift, root: v.root}
}

func (v PersistentVector) tailoff() int {
	if v.cnt < 32 {
		return 0
	}
	return ((v.cnt - 1) >> 5) << 5
}

func (v PersistentVector) arrayFor(i int) []interface{} {
	if i >= 0 && i < v.cnt {
		if i >= v.tailoff() {
			return v.tail
		}
		var node PersistentVector_Node = v.root
		for level := v.shift; level > 0; level -= 5 {
			node = node.array[(i>>level)&0x01f].(PersistentVector_Node)
		}
		return node.array
	}
	panic("index out of range")
}

func (v PersistentVector) nth(i int) interface{} {
	var node []interface{} = v.arrayFor(i)
	return node[i&0x01f]
}

func (v PersistentVector) count() int {
	return v.cnt
}

func (v PersistentVector) length() int {
	return v.cnt
}

func (v PersistentVector) cons(val interface{}) IPersistentVector {
	//i := v.cnt
	if v.cnt-v.tailoff() < 32 {
		var newTail []interface{} = make([]interface{}, len(v.tail)+1)
		copy(newTail, v.tail)
		newTail[len(v.tail)] = val
		return PersistentVector{v.cnt + 1, v.shift, v.root, newTail}
	}
	var newroot PersistentVector_Node
	tailnode := PersistentVector_Node{v.tail}
	newshift := v.shift
	if (v.cnt >> 5) > (1 << v.shift) {
		newroot = PersistentVector_Node{array: make([]interface{}, 0, 8)}
		newroot.array = append(newroot.array, v.root, newPath(v.shift, tailnode))
		newshift += 5
	} else {
		newroot = v.pushTail(v.shift, v.root, tailnode)
	}
	return PersistentVector{v.cnt + 1, newshift, newroot, []interface{}{val}}
	//
}

func (v PersistentVector) pushTail(level uint, parent PersistentVector_Node, tailnode PersistentVector_Node) PersistentVector_Node {
	var subidx = ((v.cnt - 1) >> level) & 0x01f
	var ret = PersistentVector_Node{parent.array}
	var nodeToInsert PersistentVector_Node
	if level == 5 {
		nodeToInsert = tailnode
	} else {
		var child = parent.array[subidx].(PersistentVector_Node)
		if child.array != nil { //
			nodeToInsert = v.pushTail(level-5, child, tailnode)
		} else {
			nodeToInsert = newPath(level-5, tailnode)
		}
	}
	//ret.array[subidx] = nodeToInsert
	ret.array = append(ret.array, nodeToInsert)
	return ret
	//
}

func (v PersistentVector) empty() IPersistentCollection {
	return PersistentVector_EMPTY
}

func (v PersistentVector) equiv(obj interface{}) bool {
	switch obj.(type) {
	case List:
		if len(obj.(List)) != v.count() {
			return false
		}
		for i, i2 := 0, obj.(List).iterator(); i2.hasNext(); i++ {
			if v.nth(i) != i2.next() {
				return false
			}
		}
		return true
	case IPersistentVector:
		ma := obj.(IPersistentVector)
		if ma.count() != v.count() {
			return false
		}
		for i := 0; i < v.count(); i++ {
			if v.nth(i) != ma.nth(i) {
				return false
			}
		}
		return true
	case Sequential:
		var ms ISeq = seq(obj)
		for i := 0; i < v.count(); i, ms = i+1, ms.next() {
			if ms == nil || v.nth(i) != ms.first() {
				return false
			}
		}
		if ms != nil {
			return false
		}
	}
	return false
}

func (v PersistentVector) chunkedSeq() IChunkedSeq {
	if v.count() == 0 {
		return nil
	}
	return ChunkedSeq{vec: v, i: 0, offset: 0, node: v.arrayFor(0)}
}

func (v PersistentVector) seq() ISeq {
	return v.chunkedSeq()
}

func (t TransientVector) persistent() PersistentVector {
	var trimmedTail []interface{} = t.tail
	return PersistentVector{t.cnt, t.shift, t.root, trimmedTail}
}

func (t TransientVector) conj(val interface{}) TransientVector {
	i := t.cnt
	//room is tail? = i is not multiples of 32?
	if i-t.tailoff() < 32 {
		//t.tail[i & 0x01f] = val
		t.tail = append(t.tail, val)
		t.cnt++
		return t
	}
	var newroot PersistentVector_Node
	var tailnode PersistentVector_Node = PersistentVector_Node{array: t.tail}
	t.tail = make([]interface{}, 0, 8)
	t.tail = append(t.tail, val)
	var newshift uint = t.shift
	if (t.cnt >> 5) > (1 << t.shift) {
		newroot = PersistentVector_Node{array: make([]interface{}, 0, 8)}
		newroot.array[0] = t.root
		newroot.array[1] = newPath(t.shift, tailnode)
		newshift += 5
	} else {
		newroot = t.pushTail(t.shift, t.root, tailnode)
	}
	t.root = newroot
	t.shift = newshift
	t.cnt++
	return t
}

func (t TransientVector) pushTail(level uint, parent PersistentVector_Node, tailnode PersistentVector_Node) PersistentVector_Node {
	var subidx = ((t.cnt - 1) >> level) & 0x01f
	var ret PersistentVector_Node = parent
	var nodeToInsert PersistentVector_Node
	if level == 5 {
		nodeToInsert = tailnode
	} else {
		if child, ok := parent.array[subidx].(PersistentVector_Node); ok {
			nodeToInsert = newPath(level-5, tailnode)
		} else {
			//panic("Unknown Error")
			nodeToInsert = t.pushTail(level-5, child, tailnode)
		}
	}
	//ret.array[subidx] = nodeToInsert
	ret.array = append(ret.array, nodeToInsert)
	return ret
}

func newPath(level uint, node PersistentVector_Node) PersistentVector_Node {
	if level == 0 {
		return node
	}
	var ret = PersistentVector_Node{array: make([]interface{}, 0, 8)}
	ret.array[0] = newPath(level-5, node)
	return ret
}

func (t TransientVector) tailoff() int {
	if t.cnt < 32 {
		return 0
	}
	return ((t.cnt - 1) >> 5) << 5
}

func (c ChunkedSeq) chunkedFirst() IChunk {
	return ArrayChunk{c.node, c.offset, len(c.node)}
}

func (c ChunkedSeq) chunkedNext() ISeq {
	if c.i+len(c.node) < c.vec.cnt {
		return ChunkedSeq{vec: c.vec, i: len(c.node), offset: 0}
	}
	return nil
}

func (c ChunkedSeq) chunkedMore() ISeq {
	var s ISeq = c.chunkedNext()
	if s == nil {
		return PersistentList{}
	}
	return s
}

func (c ChunkedSeq) first() interface{} {
	return c.node[c.offset]
}

func (c ChunkedSeq) next() ISeq {
	if c.offset + 1 < len(c.node) {
		return ChunkedSeq{vec: c.vec, node: c.node, i: c.i, offset: c.offset + 1}
	}
	return c.chunkedNext()
}

func (c ChunkedSeq) count() int {
	return c.vec.cnt - (c.i + c.offset)
}
