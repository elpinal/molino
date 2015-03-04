package lang

type Cons struct {
	ASeq
	_first interface{}
	_more  ISeq
}

func (c Cons) first() interface{} {
	return c._first
}

func (c Cons) next() ISeq {
	return c.more().seq()
}

func (c Cons) more() ISeq {
	if c._more == nil {
		return EmptyList{}
	}
	return c._more
}

func (c Cons) count() int {
	return 1 + c._more.count()
}

func (c Cons) seq() ISeq {
	return c
}
