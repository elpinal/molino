package molino

type SeqIterator struct {
	seq   interface{}
	_next interface{}
}

var START interface{}

func (s *SeqIterator) hasNext() bool {
	if s.seq == START {
		s.seq = nil
		s._next = seq(s._next)
	} else if s.seq == s._next {
		s._next = next(s.seq)
	}
	return s._next != nil
}

func (s *SeqIterator) next() interface{} {
	if !s.hasNext() {
		panic("Index out of range")
	}
	s.seq = s._next
	return first(s._next)
}
