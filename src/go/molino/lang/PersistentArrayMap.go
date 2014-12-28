package lang

type PersistentArrayMap struct {
	array []interface{}
}

var HASHTABLE_THRESHOLD int = 16

func (a PersistentArrayMap) assoc(key, val interface{}) IPersistentMap {
	i := a.indexOf(key)
	var newarray []interface{}
	var _, _ = i, newarray
	//
	return PersistentArrayMap{}
}

func (a PersistentArrayMap) indexOf(key interface{}) int {
	//
	return 0
}

/*
func (a PersistentArrayMap) iterator() Iterator {
	//
	return IteratorSeq{}
}
*/
