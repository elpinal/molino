package molino

type ILookup interface {
	valAt(interface{}) interface{}
}
