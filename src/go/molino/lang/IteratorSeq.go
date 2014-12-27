package lang

type IteratorSeq struct {
	ASeq
	iter *Iterator
	state State
}

type State struct {
	val   interface{}
	_rest interface{}
}

func (s IteratorSeq) create(iter *Iterator) (IteratorSeq, bool) {
	if len(iter.slice) <= iter.n {
		return IteratorSeq{}, false
	}
	state := State{}
	return IteratorSeq{iter: iter, state: State{val: state, _rest: state}}, true
}

func (s IteratorSeq) first() interface{} {
	//if s.state.val != nil {
	s.state.val = s.iter.next()
	//}
	return s.state.val
}

func (s IteratorSeq) next() ISeq {
	s.first()
	var ok bool
	s.state._rest, ok = s.create(s.iter)
	if !ok {
		return nil
	}
	return s.state._rest.(ISeq)
}
