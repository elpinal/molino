package lang

type IFn interface {
	invoke(...interface{}) interface{}
}

type ReaderFn interface {
	invoke(*Reader, rune) interface{}
}

