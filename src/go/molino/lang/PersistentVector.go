package lang

type PersistentVector struct {
	cnt   int
	shift int
}

type PersistentVector_Node struct {
	//edit
	array [32]interface{}
}

var PersistentVector_EMPTY_NODE = PersistentVector_Node{array: [32]interface{}{}}

var PersistentVector_EMPTY = PersistentVector{cnt: 0, shift: 5}

func (v PersistentVector) create(items ISeq) PersistentVector {
	var ret = PersistentVector_EMPTY
	_ = ret
	return PersistentVector{}
}

func (v PersistentVector) length() int {
	return v.cnt
}
