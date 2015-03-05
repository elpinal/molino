package molino

type ArrayChunk struct {
	array []interface{}
	off   int
	end   int
}

func (c ArrayChunk) dropFirst() IChunk {
	if c.off == c.end {
		panic("dropFirst of empty chunk")
	}
	return ArrayChunk{array: c.array, off: c.off + 1, end: c.end}
}
