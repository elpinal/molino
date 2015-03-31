package molino

type Iterable interface {
	iterator() Iterator
}
