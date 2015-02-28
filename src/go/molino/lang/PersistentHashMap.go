package lang

import (
	"fmt"
	"reflect"
	"strconv"
)

type PersistentHashMap struct {
	count    int
	root     INode
	hasNil   bool
	nilValue interface{}
}

type TransientHashMap struct {
	cnt      int
	root     INode
	hasNil   bool
	nilValue interface{}
	leafFlag Box
	edit     bool
}

type INode interface {
	assoc(int, int, interface{}, interface{}, *Box) INode
	assoc6(bool, int, int, interface{}, interface{}, *Box) INode
	findEntry(int, int, interface{}) IMapEntry
	find(int, int, interface{}, interface{}) interface{}
	nodeSeq() ISeq
}

type BitmapIndexedNode struct {
	bitmap int
	array  []interface{}
	edit   bool
}

type NodeSeq struct {
	ASeq
	array []interface{}
	i     int
	s     ISeq
}

func hash(k interface{}) int {
	if k == nil {
		return 0
	}
	switch k.(type) {
	case IHashEq:
		return k.(IHashEq).hasheq()
	case int:
		return hashInt(k.(int))
	case int64:
		return hashInt(int(k.(int64)))
	}
	n, err := strconv.ParseInt(fmt.Sprintf("%p", &k), 0, 0)
	if err != nil {
		panic(err)
	}
	return int(n)
	//panic(fmt.Sprintf("Cannot create hash from %T", k))
}

func (h PersistentHashMap) create(init ...interface{}) PersistentHashMap {
	var ret ITransientMap = PersistentHashMap{}.asTransient()
	for i := 0; i < len(init); i += 2 {
		ret = ret.assoc(init[i], init[i+1])
	}
	return ret.persistent().(PersistentHashMap)
}

func (h PersistentHashMap) createWithCheck(init []interface{}) PersistentHashMap {
	var ret ITransientMap = PersistentHashMap{}.asTransient()
	for i := 0; i < len(init); i += 2 {
		ret = ret.assoc(init[i], init[i+1])
		if ret.count() != i/2+1 {
			panic(fmt.Sprintf("Duplicate key: %s", init[i]))
		}
	}
	return ret.persistent().(PersistentHashMap)
}

func (h PersistentHashMap) entryAt(key interface{}) IMapEntry {
	if key == nil {
		//
		return nil
	}
	//
	return h.root.findEntry(0, hash(key), key)
}

func (h PersistentHashMap) assoc(key, val interface{}) Associative {
	if key == nil {
		if h.hasNil {
			if val == h.nilValue {
				return h
			}
			return PersistentHashMap{count: h.count, root: h.root, hasNil: true, nilValue: val}
		}
		return PersistentHashMap{count: h.count + 1, root: h.root, hasNil: true, nilValue: val}
	}
	var addedLeaf Box = Box{}
	var newroot INode
	if h.root == nil {
		newroot = BitmapIndexedNode{}.assoc(0, hash(key), key, val, &addedLeaf)
	} else {
		newroot = h.root.assoc(0, hash(key), key, val, &addedLeaf)
	}
	if newroot == h.root {
		return h
	}
	if addedLeaf.val == nil {
		return PersistentHashMap{count: h.count, root: newroot, hasNil: h.hasNil, nilValue: h.nilValue}
	}
	return PersistentHashMap{count: h.count + 1, root: newroot, hasNil: h.hasNil, nilValue: h.nilValue}
}

func (h PersistentHashMap) valAt(key interface{}) interface{} {
	if key == nil {
		if h.hasNil {
			return h.nilValue
		}
		return nil
	}
	if h.root != nil {
		return h.root.find(0, hash(key), key, nil)
	}
	return nil
}

func (h PersistentHashMap) seq() ISeq {
	var s ISeq = nil
	if h.root != nil {
		s = h.root.nodeSeq()
	}
	if h.hasNil {
		return Cons{_first: MapEntry{nil, h.nilValue}, _more: s}
	}
	return s
}

func (h PersistentHashMap) asTransient() TransientHashMap {
	return TransientHashMap{root: h.root, cnt: h.count}
}

func (t TransientHashMap) assoc(key, val interface{}) ITransientMap {
	return t.doAssoc(key, val)
}

func (t TransientHashMap) doAssoc(key, val interface{}) ITransientMap {
	if key == nil {
		if t.nilValue != val {
			t.nilValue = val
		}
		if !t.hasNil {
			t.cnt++
			t.hasNil = true
		}
		return t
	}
	t.leafFlag.val = nil
	var n INode
	if t.root == nil {
		n = BitmapIndexedNode{}.assoc6(t.edit, 0, hash(key), key, val, &t.leafFlag)
	} else {
		n = t.root.assoc6(t.edit, 0, hash(key), key, val, &t.leafFlag)
	}
	//
	if !reflect.DeepEqual(n, t.root) {
		t.root = n
	}
	if t.leafFlag.val != nil {
		t.cnt++
	}
	return t
}

func (t TransientHashMap) persistent() IPersistentMap {
	return t.doPersistent()
}

func (t TransientHashMap) doPersistent() IPersistentMap {
	return PersistentHashMap{root: t.root, count: t.cnt}
}

func (t TransientHashMap) count() int {
	return t.cnt
}

func (b BitmapIndexedNode) index(bit int) int {
	return bitCount(b.bitmap & (bit - 1))
}

func (b BitmapIndexedNode) assoc(shift int, hash int, key interface{}, val interface{}, addedLeaf *Box) INode {
	bit := bitpos(hash, shift)
	idx := b.index(bit)
	if (b.bitmap & bit) != 0 {
		keyOrNil := b.array[2*idx]
		valOrNode := b.array[2*idx+1]
		if keyOrNil == nil {
			var n INode = valOrNode.(INode).assoc(shift+5, hash, key, val, addedLeaf)
			if reflect.DeepEqual(n, valOrNode) {
				return b
			}
			return BitmapIndexedNode{bitmap: b.bitmap, array: cloneAndSet(b.array, 2*idx+1, n)}
		}
		if key == keyOrNil {
			if val == valOrNode {
				return b
			}
			return BitmapIndexedNode{bitmap: b.bitmap, array: cloneAndSet(b.array, 2*idx+1, val)}
		}
		addedLeaf.val = addedLeaf
		//
		return BitmapIndexedNode{bitmap: b.bitmap, array: cloneAndSet(b.array, 2*idx, nil, 2*idx+1, createNode(shift+5, keyOrNil, valOrNode, hash, key, val))}
	}
	n := bitCount(b.bitmap)
	if n >= 16 {
		panic("FIXME:")
	}

	var newArray []interface{} = make([]interface{}, 0, 2*(n+1))
	for i := 0; i < 2*idx; i++ {
		newArray = append(newArray, b.array[i])
	}
	newArray = append(newArray, key)
	addedLeaf.val = addedLeaf
	newArray = append(newArray, val)
	for i := 2 * idx; i < 2*n; i++ {
		newArray = append(newArray, b.array[i])
	}
	return BitmapIndexedNode{bitmap: b.bitmap | bit, array: newArray}
}

func (b BitmapIndexedNode) ensureEditable(edit bool) BitmapIndexedNode {
	if b.edit == edit {
		return b
	}
	n := bitCount(b.bitmap)
	var newArray []interface{}
	if n >= 0 {
		newArray = make([]interface{}, 2*(n+1))
	} else {
		newArray = make([]interface{}, 4)
	}
	copy(newArray, b.array)
	return BitmapIndexedNode{edit: edit, bitmap: b.bitmap, array: newArray}
}

func (b BitmapIndexedNode) editAndSet(edit bool, i int, a interface{}) BitmapIndexedNode {
	var editable BitmapIndexedNode = b.ensureEditable(edit)
	editable.array[i] = a
	return editable
}

func (b BitmapIndexedNode) editAndSet5(edit bool, i int, a interface{}, j int, c interface{}) BitmapIndexedNode {
	var editable BitmapIndexedNode = b.ensureEditable(edit)
	editable.array[i] = a
	editable.array[j] = c
	return editable
}

func (b BitmapIndexedNode) assoc6(edit bool, shift int, hash int, key interface{}, val interface{}, addedLeaf *Box) INode {
	bit := bitpos(hash, shift)
	idx := b.index(bit)
	if (b.bitmap & bit) != 0 {
		keyOrNil := b.array[2*idx]
		valOrNode := b.array[2*idx+1]
		if keyOrNil == nil {
			var n INode = valOrNode.(INode).assoc6(edit, shift+5, hash, key, val, addedLeaf)
			if reflect.DeepEqual(n, valOrNode) {
				return b
			}
			return b.editAndSet(edit, 2*idx+1, n)
		}
		if key == keyOrNil {
			if val == valOrNode {
				return b
			}
			return b.editAndSet(edit, 2*idx+1, val)
		}
		addedLeaf.val = addedLeaf
		//
		return b.editAndSet5(edit, 2*idx, nil, 2*idx+1, createNode(edit, shift+5, keyOrNil, valOrNode, hash, key, val))
	}
	n := bitCount(b.bitmap)
	if n*2 < len(b.array) {
		addedLeaf.val = addedLeaf
		var editable BitmapIndexedNode = b.ensureEditable(edit)
		_ = editable
		panic("FIXME:")
		//
	}
	if n >= 16 {
		panic("FIXME:")
	}

	var newArray []interface{} = make([]interface{}, 0, 2*(n+4))
	for i := 0; i < 2*idx; i++ {
		newArray = append(newArray, b.array[i])
	}
	newArray = append(newArray, key)
	addedLeaf.val = addedLeaf
	newArray = append(newArray, val)
	for i := 2 * idx; i < 2*n; i++ {
		newArray = append(newArray, b.array[i])
	}
	var editable BitmapIndexedNode = b.ensureEditable(edit)
	editable.array = newArray
	editable.bitmap |= bit
	return editable
}

func (b BitmapIndexedNode) findEntry(shift int, hash int, key interface{}) IMapEntry {
	bit := bitpos(hash, shift)
	if (b.bitmap & bit) == 0 {
		return nil
	}
	idx := b.index(bit)
	keyOrNil := b.array[2*idx]
	valOrNode := b.array[2*idx+1]
	if keyOrNil == nil {
		return valOrNode.(INode).findEntry(shift+5, hash, key)
	}
	if key == keyOrNil {
		return MapEntry{keyOrNil, valOrNode}
	}
	return nil
}

func (b BitmapIndexedNode) find(shift int, hash int, key, notFound interface{}) interface{} {
	bit := bitpos(hash, shift)
	if (b.bitmap & bit) == 0 {
		return nil
	}
	idx := b.index(bit)
	keyOrNil := b.array[2*idx]
	valOrNode := b.array[2*idx+1]
	if keyOrNil == nil {
		return valOrNode.(INode).find(shift+5, hash, key, notFound)
	}
	if key == keyOrNil {
		return valOrNode
	}
	return notFound
}

func (b BitmapIndexedNode) nodeSeq() ISeq {
	return NodeSeq{}.create(b.array, 0, nil)
}

func (_ NodeSeq) create(array []interface{}, i int, s ISeq) ISeq {
	if s != nil {
		return NodeSeq{array: array, i: i, s: s}
	}
	for j := i; j < len(array); j += 2 {
		if array[j] != nil {
			return NodeSeq{array: array, i: j}
		}
		var node INode = array[j+1].(INode)
		if node != nil {
			var nodeSeq ISeq = node.nodeSeq()
			if nodeSeq != nil {
				return NodeSeq{array: array, i: j + 2, s: nodeSeq}
			}
		}
	}
	return nil
}

func (n NodeSeq) first() interface{} {
	if n.s != nil {
		return n.s.first()
	}
	return MapEntry{_key: n.array[n.i], _val: n.array[n.i+1]}
}

func (n NodeSeq) next() ISeq {
	if n.s != nil {
		return n.create(n.array, n.i, n.s.next())
	}
	return n.create(n.array, n.i+2, nil)
}

func mask(hash int, shift int) uint {
	if shift < 0 {
		panic("Stupid shift")
	}
	return uint((hash >> uint(shift)) & 0x01f)
}

func bitpos(hash int, shift int) int {
	return 1 << mask(hash, shift)
}

func bitCount(i int) int {
	i = i - ((i >> 1) & 0x55555555)
	i = (i & 0x33333333) + ((i >> 2) & 0x33333333)
	i = (i + (i >> 4)) & 0x0f0f0f0f
	i = i + (i >> 8)
	i = i + (i >> 16)
	return i & 0x3f
}

func createNode(edit bool, shift int, key1 interface{}, val1 interface{}, key2hash int, key2 interface{}, val2 interface{}) INode {
	key1hash := hash(key1)
	if key1hash == key2hash {
		//
		panic("FIXME: Hash collision: " + fmt.Sprint(key1, key2, key1hash, key2hash))
	}
	//
	var addedLeaf Box = Box{nil}
	return BitmapIndexedNode{}.assoc6(edit, shift, key1hash, key1, val1, &addedLeaf).assoc6(edit, shift, key2hash, key2, val2, &addedLeaf)
}
