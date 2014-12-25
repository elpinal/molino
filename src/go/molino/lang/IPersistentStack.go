package lang

type IPersistentStack interface {
	IPersistentCollection
	peek() interface{}
	pop() IPersistentStack
}
