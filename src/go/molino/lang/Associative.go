package lang

type Associative interface {
	//IPersistentCollection
	Seqable
	ILookup
	entryAt(interface{}) IMapEntry
	assoc(interface{}, interface{}) Associative
}
