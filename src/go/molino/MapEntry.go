package molino

type MapEntry struct {
	_key interface{}
	_val interface{}
}

func (e MapEntry) key() interface{} {
	return e._key
}

func (e MapEntry) val() interface{} {
	return e._val
}
