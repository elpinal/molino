package lang

type IPersistentMap interface {
	Iterable
	assoc(interface{}, interface{}) IPersistentMap
}
