package molino

type SeqIterator struct {
	seq   interface{}
	nexts interface{}
}

var START interface{}

func (s *SeqIterator) hasNext() bool {
	if s.seq == START {
		s.seq = nil
		s.nexts = seq(s.nexts)
	} else if s.seq == s.nexts {
		s.nexts = next(s.seq)
	}
	return s.nexts != nil
}

func (s *SeqIterator) next() interface{} {
	if !s.hasNext() {
		panic("Index out of range")
	}
	s.seq = s.nexts
	return first(s.nexts)
}
