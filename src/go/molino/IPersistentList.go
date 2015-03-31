package molino

type IPersistentList interface {
	Sequential
	IPersistentStack
	cons(interface{}) ISeq
}
