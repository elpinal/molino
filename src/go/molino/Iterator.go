package molino

type Iterator interface {
	hasNext() bool
	next() interface{}
}
