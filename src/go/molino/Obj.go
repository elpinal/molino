package molino

type Obj struct {
	_meta IPersistentMap
}

func (o Obj) meta() IPersistentMap {
	if o._meta == nil {
		return PersistentHashMap{}
	}
	return o._meta
}
