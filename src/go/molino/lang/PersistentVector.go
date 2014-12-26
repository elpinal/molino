package lang

type PersistentVector struct {
	cnt   int
	shift uint
	root  PersistentVector_Node
	tail  [32]interface{}
}

type PersistentVector_Node struct {
	//edit
	array [32]interface{}
}

type TransientVector struct {
	cnt   int
	shift uint
	root  PersistentVector_Node
	tail  [32]interface{}
}

var PersistentVector_EMPTY_NODE = PersistentVector_Node{array: [32]interface{}{}}

var PersistentVector_EMPTY = PersistentVector{cnt: 0, shift: 5, root: PersistentVector_EMPTY_NODE}

func (v PersistentVector) create(items ISeq) PersistentVector {
	var ret = PersistentVector_EMPTY.asTransient()
	_ = ret
	return PersistentVector{}
}

func (v PersistentVector) asTransient() TransientVector {
	return TransientVector{cnt: v.cnt, shift: v.shift, root: v.root}
}

func (v PersistentVector) length() int {
	return v.cnt
}

func newPath(level uint, node PersistentVector_Node) PersistentVector_Node {
	if level == 0 {
		return node
	}
	var ret = PersistentVector_Node{array: [32]interface{}{}}
	ret.array[0] = newPath(level - 5, node)
	return ret
}

func (t TransientVector) conj(val interface{}) TransientVector {
	i := t.cnt
	//room is tail? = i is not multiples of 32?
	if i - t.tailoff() < 32 {
		t.tail[i & 0x01f] = val
		t.cnt++
		return t
	}
	var newroot PersistentVector_Node
	var tailnode PersistentVector_Node = PersistentVector_Node{array: t.tail}
	t.tail = [32]interface{}{}
	t.tail[0] = val
	var newshift uint = t.shift
	if (t.cnt >> 5) > (1 << t.shift) {
		newroot = PersistentVector_Node{array: [32]interface{}{}}
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
			nodeToInsert = //t.pushTail()
		} else {
			panic("Unknown Error")
		}
	}
	ret.array[subidx] = nodeToInsert
	return ret
}

func (t TransientVector) tailoff() int {
	if t.cnt < 32 {
		return 0
	}
	return ((t.cnt - 1) >> 5) << 5
}
