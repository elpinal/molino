package lang

type IPersistentCollection interface {
	Seqable
	count() int
	cons(interface{}) IPersistentCollection
	empty() IPersistentCollection
	equiv(interface{}) bool
}
