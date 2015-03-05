package molino

type AReference struct {
	_meta IPersistentMap
}

func (r AReference) meta() IPersistentMap {
	if r._meta == nil {
		return PersistentHashMap{}
	}
	return r._meta
}

func (r *AReference) resetMeta(m IPersistentMap) IPersistentMap {
	r._meta = m
	return m
}
