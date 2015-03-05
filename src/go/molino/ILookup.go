package lang

type ILookup interface {
	valAt(interface{}) interface{}
}
