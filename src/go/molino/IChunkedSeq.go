package molino

type IChunkedSeq interface {
	ISeq
	chunkedFirst() IChunk
	chunkedNext() ISeq
	chunkedMore() ISeq
}
