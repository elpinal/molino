package molino

type ISeq interface {
	IPersistentCollection
	first() interface{}
	next() ISeq
	more() ISeq
	cons(interface{}) ISeq
}
