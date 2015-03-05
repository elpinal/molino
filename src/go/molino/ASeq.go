package molino

type ASeq struct {
	Obj
	_hash   int
	_hasheq int
}

func (a ASeq) count() int {
	i := 1
	for s := a.next(); s != nil; s, i = s.next(), i+1 {
		if _, ok := s.(Counted); ok {
			return i + s.count()
		}
	}
	return i
}

func (a ASeq) seq() ISeq {
	return a
}

func (a ASeq) cons(o interface{}) ISeq {
	return Cons{_first: o, _more: a}
}

func (a ASeq) empty() IPersistentCollection {
	return EmptyList{}
}

func (a ASeq) equiv(obj interface{}) bool {
	switch obj.(type) {
	case Sequential, List:
		var ms ISeq = seq(obj)
		for s := a.seq(); s != nil; s, ms = s.next(), ms.next() {
			if ms == nil || s.first() != ms.first() {
				return false
			}
		}
		return ms == nil
	}
	return false
}

func (a ASeq) first() interface{} {
	return nil
}

func (a ASeq) more() ISeq {
	var s ISeq = a.next()
	if s == nil {
		return EmptyList{}
	}
	return s
}
func (a ASeq) next() ISeq {
	return EmptyList{}
}
