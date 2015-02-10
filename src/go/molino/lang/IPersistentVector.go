package lang

type IPersistentVector interface {
	Sequential
	count() int
	cons(interface{}) IPersistentVector
	length() int
	nth(int) interface{}
}
