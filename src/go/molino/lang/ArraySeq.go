package lang

type ArraySeq struct {
	ASeq
	array List
	i     int
}

func (a ArraySeq) first() interface{} {
	if a.array != nil {
		return a.array[a.i]
	}
	return nil
}

func (a ArraySeq) next() ISeq {
	if a.array != nil && a.i + 1 < len(a.array) {
		return ArraySeq{array: a.array, i: a.i}
	}
	return nil
}
