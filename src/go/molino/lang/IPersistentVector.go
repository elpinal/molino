package lang

type IPersistentVector interface {
	Sequential
	length() int
}
