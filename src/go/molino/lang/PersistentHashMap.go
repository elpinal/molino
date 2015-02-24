package lang

import (
	"fmt"
)

type PersistentHashMap struct {
	count    int
	root     INode
	hasNil   bool
	nilValue interface{}
}

type TransientHashMap struct {
	count    int
	root     INode
	hasNil   bool
	nilValue interface{}
	leafFlag Box
	edit     chan bool
}

type INode interface {
	assoc(int, int, interface{}, interface{}, Box) INode
	assoc6(chan bool, int, uint, interface{}, interface{}, Box) INode
	find(int, uint, interface{}) IMapEntry
}

type BitmapIndexedNode struct {
	bitmap int
	array  []interface{}
	edit   chan bool
}


func hash(k interface{}) uint {
	fmt.Printf("2 %s\n", k)
	if k == nil {
		return 0
	}
	switch k.(type) {
	case int:
		return hashInt(k.(int))
	}
	panic("Cannot create hash")
}

func (h PersistentHashMap) createWithCheck(init []interface{}) PersistentHashMap {
	var ret ITransientMap = PersistentHashMap{}.asTransient()
	for i := 0; i < len(init); i += 2 {
		ret = ret.assoc(init[i], init[i + 1])
	}
	return ret.persistent().(PersistentHashMap)
}

func (h PersistentHashMap) entryAt(key interface{}) IMapEntry {
	if key == nil {
		//
		return nil
	}
	//
	return h.root.find(0, hash(key), key)
}

func (h PersistentHashMap) assoc(key, val interface{}) Associative {
	//var ret ITransientMap = PersistentHashMap{}.asTransient()
	//
	return PersistentHashMap{}
}

func (h PersistentHashMap) valAt(key interface{}) interface{} {
	if key == nil {
		if h.hasNil {
			return h.nilValue
		}
		return nil
	}
	if h.root != nil {
		h.root.find(0, hash(key), key)
	}
	return nil
}

func (h PersistentHashMap) asTransient() TransientHashMap {
	return TransientHashMap{root: h.root, count: h.count}
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
			t.count++
			t.hasNil = true
		}
		return t
	}
	t.leafFlag.val = nil
	var n INode
	if t.root == nil {
		n = BitmapIndexedNode{}.assoc6(t.edit, 0, hash(key), key, val, t.leafFlag)
	} else {
		n = t.root.assoc6(t.edit, 0, hash(key), key, val, t.leafFlag)
	}
	//
	//if n != t.root {
		t.root = n
	//}
	if t.leafFlag.val != nil {
		t.count++
	}
	return t
}

func (t TransientHashMap) persistent() IPersistentMap {
	return t.doPersistent()
}

func (t TransientHashMap) doPersistent() IPersistentMap {
	return PersistentHashMap{root: t.root, count: t.count}
}


func (b BitmapIndexedNode) index(bit int) int {
	return bitCount(b.bitmap & (bit - 1))
}

func (b BitmapIndexedNode) assoc(shift int, hash int, key interface{}, val interface{}, addedLeaf Box) INode {
	//
	return BitmapIndexedNode{}
}

func (b BitmapIndexedNode) ensureEditable(edit chan bool) BitmapIndexedNode {
	if b.edit == edit {
		return b
	}
	n := bitCount(b.bitmap)
	var newArray []interface{}
	if n >= 0 {
		newArray = make([]interface{}, 2 * (n + 1))
	} else {
		newArray = make([]interface{}, 4)
	}
	copy(newArray, b.array)
	return BitmapIndexedNode{edit: edit, bitmap: b.bitmap, array: newArray}
}

func (b BitmapIndexedNode) editAndSet(edit chan bool, i int, a interface{}) BitmapIndexedNode {
	var editable BitmapIndexedNode = b.ensureEditable(edit)
	editable.array[i] = a
	return editable
}

func (b BitmapIndexedNode) editAndSet5(edit chan bool, i int, a interface{}, j int, c interface{}) BitmapIndexedNode {
	var editable BitmapIndexedNode = b.ensureEditable(edit)
	editable.array[i] = a
	editable.array[j] = c
	return editable
}

func (b BitmapIndexedNode) assoc6(edit chan bool, shift int, hash uint, key interface{}, val interface{}, addedLeaf Box) INode {
	bit := bitpos(hash, shift)
	idx := b.index(bit)
	if (b.bitmap & bit) != 0 {
		keyOrNil := b.array[2 * idx]
		valOrNode := b.array[2 * idx + 1]
		if keyOrNil == nil {
			var n INode = valOrNode.(INode).assoc6(edit, shift + 5, hash, key, val, addedLeaf)
			if n == valOrNode {
				return b
			}
			return b.editAndSet(edit, 2 * idx + 1, n)
		}
		if key == keyOrNil {
			if val == valOrNode {
				return b
			}
			return b.editAndSet(edit, 2 * idx + 1, val)
		}
		addedLeaf.val = addedLeaf
		//
		return b.editAndSet5(edit, 2 * idx, nil, 2 * idx + 1, createNode(edit, shift + 5, keyOrNil, valOrNode, hash, key, val))
	}
	n := bitCount(b.bitmap)
	if n * 2 < len(b.array) {
		addedLeaf.val = addedLeaf
		var editable BitmapIndexedNode = b.ensureEditable(edit)
		_ = editable
		panic("FIXME:")
		//
	}
	if n >= 16 {
		panic("FIXME:")
	}

	fmt.Println(1, n, idx)
	var newArray []interface{} = make([]interface{}, 0, 2*(n+4))
	for i := 0; i < 2*idx; i++ {
		newArray = append(newArray, b.array[i])
	}
	newArray = append(newArray, key)
	addedLeaf.val = addedLeaf
	newArray = append(newArray, val)
	for i := 2*idx; i < 2*(n-idx); i++ {
		newArray = append(newArray, b.array[1])
	}
	var editable BitmapIndexedNode = b.ensureEditable(edit)
	editable.array = newArray
	editable.bitmap |= bit
	return editable
}

func (b BitmapIndexedNode) find(shift int, hash uint, key interface{}) IMapEntry {
	bit := bitpos(hash, shift)
	if (b.bitmap & bit) == 0 {
		return nil
	}
	idx := b.index(bit)
	keyOrNil := b.array[2*idx]
	valOrNode := b.array[2*idx+1]
	if keyOrNil == nil {
		return valOrNode.(INode).find(shift + 5, hash, key)
	}
	if key == keyOrNil {
		return MapEntry{keyOrNil, valOrNode}
	}
	return nil
}

func mask(hash uint, shift int) uint {
	if shift < 0 {
		panic("Stupid shift")
	}
	return uint((hash >> uint(shift)) & 0x01f)
}

func bitpos(hash uint, shift int) int {
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

func createNode(edit chan bool, shift int, key1 interface{}, val1 interface{}, key2hash uint, key2 interface{}, val2 interface{}) INode {
	key1hash := hash(key1)
	if key1hash == key2hash {
		//
		panic("FIXME: Hash collision")
	}
	//
	var addedLeaf Box = Box{nil}
	return BitmapIndexedNode{}.assoc6(edit, shift, key1hash, key1, val1, addedLeaf).assoc6(edit, shift, key2hash, key2, val2, addedLeaf)
}
