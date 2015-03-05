package lang

type IMapEntry interface {
	key() interface{}
	val() interface{}
}
