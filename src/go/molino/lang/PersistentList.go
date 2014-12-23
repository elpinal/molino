package lang

type PersistentList struct {
	_first interface{}
	_rest []interface{}
	_count int
}

