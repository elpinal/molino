package lang

type PersistentList struct {
	_first interface{}
	_rest IPersistentList
	_count int
}

type EmptyList struct {
	Obj
}


func (l PersistentList) first() interface{} {
	return l._first
}

func (l PersistentList) count() int {
	return l._count
}

func (l PersistentList) empty() IPersistentCollection {
	return EmptyList{}
}



func (e EmptyList) equals(o interface{}) bool {
	return o == nil
}

func (e EmptyList) equiv(o interface{}) bool {
	return e.equals(o)
}

func (e EmptyList) first() interface{} {
	return nil
}
func (e EmptyList) next() ISeq {
	return nil
}
func (e EmptyList) more() ISeq {
	return e
}
func (e EmptyList) cons(o interface{}) PersistentList {
	return PersistentList{}
}

func (e EmptyList) empty() IPersistentCollection {
	return e
}

func (e EmptyList) peek() interface{} {
	return nil
}

func (e EmptyList) pop() IPersistentStack {
	panic("Can't pop empty list")
}

func (e EmptyList) count() int {
	return 0
}

func (e EmptyList) seq() ISeq {
	return nil
}
