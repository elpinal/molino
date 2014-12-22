package lang

type Fn interface {
	invoke(*Reader, rune) interface{}
}

