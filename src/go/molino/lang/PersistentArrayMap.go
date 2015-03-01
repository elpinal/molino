package lang

import (
	"fmt"
)

type PersistentArrayMap struct {
	array []interface{}
}

type Seq struct {
	array []interface{}
	i     int
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

func (a PersistentArrayMap) entryAt(key interface{}) IMapEntry {
	var i int = a.indexOf(key)
	if i >= 0 {
		return MapEntry{_key: a.array[i], _val: a.array[i+1]}
	}
	return nil
}

func (a PersistentArrayMap) assoc(key, val interface{}) Associative {
	i := a.indexOf(key)
	var newArray []interface{}
	if i >= 0 {
		if a.array[i+1] == val {
			return a
		}
		newArray = make([]interface{}, len(a.array))
		copy(newArray, a.array)
		newArray[i+1] = val
	} else {
		/*
			if len(a.array) > HASHTABLE_THRESHOLD {
				return //
			}
		*/
		newArray = make([]interface{}, 2, len(a.array)+2)
		if len(a.array) > 0 {
			newArray = append(newArray, a.array...)
		}
		newArray[0] = key
		newArray[1] = val
	}
	return PersistentArrayMap{array: newArray}
}

func (a PersistentArrayMap) valAt(key interface{}) interface{} {
	var i int = a.indexOf(key)
	if i >= 0 {
		return a.array[i+1]
	}
	return nil
}

func (a PersistentArrayMap) indexOfObject(key interface{}) int {
	//
	for i := 0; i < len(a.array); i += 2 {
		if key == a.array[i] {
			return i
		}
	}
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

func (a PersistentArrayMap) seq() ISeq {
	if len(a.array) > 0 {
		return Seq{a.array, 0}
	}
	return nil
}

func (a Seq) empty() IPersistentCollection {
	return EmptyList{}
}

func (a Seq) equiv(obj interface{}) bool {
	switch obj.(type) {
	case Sequential, List:
		var ms ISeq = seq(obj)
		for s := a.seq(); s != nil; s, ms = s.next(), ms.next() {
			if ms == nil || s.first() != ms.first() {
				return false
			}
		}
	}
	return false
}

func (a Seq) more() ISeq {
	s := a.next()
	if s == nil {
		return EmptyList{}
	}
	return s
}

func (a Seq) seq() ISeq {
	return a
}

func (a Seq) first() interface{} {
	return MapEntry{a.array[a.i], a.array[a.i+1]}
}

func (a Seq) next() ISeq {
	if a.i+2 < len(a.array) {
		return Seq{a.array, a.i + 2}
	}
	return nil
}

func (a Seq) count() int {
	return (len(a.array) - 1) / 2
}

func (a Seq) cons(o interface{}) ISeq {
	return Cons{_first: o, _more: a}
}
