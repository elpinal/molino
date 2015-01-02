package lang

type PersistentHashMap struct {
	count    int
	hasNil   bool
	nilValue interface{}
}

type TransientHashMap struct {
	count    int
	hasNil   bool
	nilValue interface{}
	leafFlag Box
}

type INode interface {
	assoc(int, int, interface{}, interface{}, Box) INode
	assoc6(chan bool, int, int, interface{}, interface{}, Box) INode
}

type BitmapIndexedNode struct {
	bitmap int
	array  []interface{}
	edit   chan bool
}


func hash(k interface{}) int {
	//
	return 0
}

func (h PersistentHashMap) createWithCheck(init []interface{}) PersistentHashMap {
	var ret ITransientMap = PersistentHashMap{}.asTransient()
	for i := 0; i < len(init); i += 2 {
		ret = ret.assoc(init[i], init[i + 1])
	}
	return ret.persistent().(PersistentHashMap)
}

func (h PersistentHashMap) assoc(key, val interface{}) IPersistentMap {
	//var ret ITransientMap = PersistentHashMap{}.asTransient()
	//
	return PersistentHashMap{}
}

func (h PersistentHashMap) asTransient() TransientHashMap {
	return TransientHashMap{count: h.count}
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
	//var n INode = root.assoc(0, hash(key), key, t.leafFlag)
	//
	if t.leafFlag.val != nil {
		t.count++
	}
	return t
}

func (t TransientHashMap) persistent() IPersistentMap {
	return t.doPersistent()
}

func (t TransientHashMap) doPersistent() IPersistentMap {
	return PersistentHashMap{count: t.count}
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

func (b BitmapIndexedNode) assoc6(edit chan bool, shift int, hash int, key interface{}, val interface{}, addedLeaf Box) INode {
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
		//return b.editAndSet(edit, 2 * idx, nil, 2 * idx + 1, createNode(edit, shift + 5, keyOrNil, valOrNode, hash, key, val))
		return BitmapIndexedNode{}
	}
	n := bitCount(b.bitmap)
	if n * 2 < len(b.array) {
		addedLeaf.val = addedLeaf
		//
	}
	//
	return BitmapIndexedNode{}
}

func mask(hash, shift int) uint {
	if shift < 0 {
		panic("Stupid shift")
	}
	return uint((hash >> uint(shift)) & 0x01f)
}

func bitpos(hash, shift int) int {
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

func createNode(edit chan bool, shift int, key1 interface{}, val1 interface{}, key2hash int, key2 interface{}, val2 interface{}) INode {
	key1hash := hash(key1)
	_ = key1hash
	//
	return BitmapIndexedNode{}
}
