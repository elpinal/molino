package molino

type IPersistentCollection interface {
	Seqable
	count() int
	//cons(interface{}) IPersistentCollection
	//cons(interface{}) ISeq
	empty() IPersistentCollection
	equiv(interface{}) bool
}
