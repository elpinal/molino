package lang

type PersistentList struct {
	_first interface{}
	_rest []interface{}
	_count int
}

type EmptyList struct {
	Obj
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
func (e EmptyList) cons(o interface{}) ISeq {
	return nil
}
