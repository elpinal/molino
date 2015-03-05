package molino

type IPersistentList interface {
	Sequential
	IPersistentStack
}
