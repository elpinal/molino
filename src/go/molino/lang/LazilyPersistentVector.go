package lang

type LazilyPersistentVector struct {
}

func (_ LazilyPersistentVector) create(obj interface{}) IPersistentVector {
	/*
	if _, ok := obj.(ISeq); !ok && len(obj) <= 32 {
		//
	}
	*/
	//
	if iter, ok := obj.(Iterable); ok {
		return PersistentVector{}.create(iter)
	} else if i, ok := obj.([]interface{}); ok {
		var list List = i
		return PersistentVector{}.create(list)
	} else if _, ok := obj.(ISeq); ok {
		return PersistentVector{}.create(seq(obj))
	} else {
		return PersistentVector{}
	}
}
