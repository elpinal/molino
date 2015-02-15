package lang

type Associative interface {
	entryAt(interface{}) IMapEntry
	assoc(interface{}, interface{}) Associative
}
