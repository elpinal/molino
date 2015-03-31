package molino

type IMapEntry interface {
	key() interface{}
	val() interface{}
}
