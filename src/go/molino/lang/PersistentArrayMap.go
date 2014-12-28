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

func (a PersistentArrayMap) indexOfObject(key interface{}) int {
	//
	return -1
}

func (a PersistentArrayMap) indexOf(key interface{}) int {
	if k, ok := key.(Keyword); ok {
		for i := 0; i < len(a.array); i += 2 {
			if k == a.array[i] {
				return i
			}
		}
		return -1
	}
	return a.indexOfObject(key)
}

/*
func (a PersistentArrayMap) iterator() Iterator {
	//
	return IteratorSeq{}
}
*/
