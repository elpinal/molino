package lang

type PersistentArrayMap struct {
	array []interface{}
}

var HASHTABLE_THRESHOLD int = 16

func (a PersistentArrayMap) assoc(key, val interface{}) IPersistentMap {
	return
}
