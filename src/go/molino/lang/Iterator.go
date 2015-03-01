package lang

type Iterator interface {
	hasNext() bool
	next() interface{}
}
