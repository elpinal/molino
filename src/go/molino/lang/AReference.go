package lang

type AReference struct {
	_meta IPersistentMap
}

func (r AReference) meta() IPersistentMap {
	return r._meta
}

func (r *AReference) resetMeta(m IPersistentMap) IPersistentMap {
	r._meta = m
	return m
}
