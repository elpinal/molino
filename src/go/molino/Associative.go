package lang

type Associative interface {
	IPersistentCollection
	ILookup
	containsKey(interface{}) bool
	entryAt(interface{}) IMapEntry
	assoc(interface{}, interface{}) Associative
}
