package molino

type IPersistentStack interface {
	IPersistentCollection
	peek() interface{}
	pop() IPersistentStack
}
