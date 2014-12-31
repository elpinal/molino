package lang

type ITransientMap interface {
	assoc(interface{}, interface{}) ITransientMap
	persistent() IPersistentMap
}
