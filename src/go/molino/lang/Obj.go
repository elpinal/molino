package lang

type Obj struct {
	_meta IPersistentMap
}

func (o Obj) meta() IPersistentMap {
	return o._meta
}
