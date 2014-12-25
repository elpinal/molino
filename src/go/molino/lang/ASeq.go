package lang

type ASeq struct {
	Obj
	_hash   int
	_hasheq int
}

func (a ASeq) count() int {
	//var i int = 1
	//for s := a.next(); s != nil; s, i = s.next(), i + 1 {
	//	if {
			//
	////s := a.next()
	////return i + s.count()
	//	}
	//}
	//return i
	return -1
}

func (a ASeq) seq() ISeq {
	return a
}

func (a ASeq) empty() IPersistentCollection {
	return EmptyList{}
}

func (a ASeq) equiv(o interface{}) bool {
	//
	return false
}

func (a ASeq) first() interface{} {
	return nil
}

func (a ASeq) more() ISeq {
	return EmptyList{}
}
func (a ASeq) next() ISeq {
	return EmptyList{}
}
