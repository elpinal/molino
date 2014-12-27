package lang

type ArraySeq struct {
	ASeq
	array List
	i     int
}

func createFromObject(array interface{}) ISeq {
	if array == nil {
		return nil
	}
	switch array.(type) {
	case []interface{}:
		return ArraySeq{array: array.([]interface{}), i: 0}
	}
	return nil
}

func (a ArraySeq) first() interface{} {
	if a.array != nil {
		return a.array[a.i]
	}
	return nil
}

func (a ArraySeq) next() ISeq {
	if a.array != nil && a.i + 1 < len(a.array) {
		return ArraySeq{array: a.array, i: a.i + 1}
	}
	return nil
}
