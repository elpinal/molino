package molino

type IPersistentVector interface {
	Sequential
	count() int
	cons(interface{}) IPersistentVector
	length() int
	nth(int) interface{}
}
