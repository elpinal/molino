package lang

type Iterator struct {
	n int
	slice []interface{}
}

func (i *Iterator) next() interface{} {
	nn := i.n
	if len(i.slice) <= i.n {
		return nil
	}
	i.n++
	return i.slice[nn]
}
