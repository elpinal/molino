package lang

type LazilyPersistentVector struct {
}

func (_ LazilyPersistentVector) create(coll []interface{}) IPersistentVector {
	/*
	if _, ok := coll.(ISeq); !ok && len(coll) <= 32 {
		//
	}
	*/
	return PersistentVector.create(PersistentVector{}, seq(coll))
}
