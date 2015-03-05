package molino

type List []interface{}

func (i List) iterator() Iterator {
	return &SeqIterator{seq: START, nexts: i}
}

/*
func (i *Iterator) next() interface{} {
	nn := i.n
	if len(i.slice) <= i.n {
		return nil
	}
	i.n++
	return i.slice[nn]
}
*/
