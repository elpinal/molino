package lang

import (
	"fmt"
)

type PersistentArrayMap struct {
	array []interface{}
}

var HASHTABLE_THRESHOLD int = 16

func (a PersistentArrayMap) createWithCheck(init []interface{}) PersistentArrayMap {
	for i := 0; i < len(init); i += 2 {
		for j := i + 2; j < len(init); j += 2 {
			if init[i] == init[j] {
				panic("Duplicate key: " + fmt.Sprint(init[i]))
			}
		}
	}
	return PersistentArrayMap{array: init}
}

func (a PersistentArrayMap) assoc(key, val interface{}) IPersistentMap {
	i := a.indexOf(key)
	var newArray []interface{}
	if i >= 0 {
		if a.array[i + 1] == val {
			return a
		}
		newArray = make([]interface{}, len(a.array))
		copy(newArray, a.array)
		newArray[i + 1] = val
	} else {
		/*
		if len(a.array) > HASHTABLE_THRESHOLD {
			return //
		}
		*/
		newArray = make([]interface{}, 2, len(a.array) + 2)
		if len(a.array) > 0 {
			newArray = append(newArray, a.array...)
		}
		newArray[0] = key
		newArray[1] = val
	}
	return PersistentArrayMap{array: newArray}
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
