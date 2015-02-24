package lang

type Associative interface {
	ILookup
	entryAt(interface{}) IMapEntry
	assoc(interface{}, interface{}) Associative
}
