package lang

type IFn interface {
	invoke(...interface{}) interface{}
	applyTo(ISeq) interface{}
}

type ReaderFn interface {
	invoke(*Reader, rune) (interface{}, error)
}
