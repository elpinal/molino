package molino

type IChunk interface {
	dropFirst() IChunk
}
