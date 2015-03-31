package molino

type ITransientMap interface {
	Counted
	assoc(interface{}, interface{}) ITransientMap
	persistent() IPersistentMap
}
