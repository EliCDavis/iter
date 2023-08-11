package iter

type RangeIterator struct {
	size    int
	pointer int
}

func Range(size int) *RangeIterator {
	if size < 0 {
		panic(ErrInvalidRangeValue)
	}
	return &RangeIterator{
		size:    size,
		pointer: 0,
	}
}

func (ri *RangeIterator) Next() (item int, err error) {
	if ri.pointer >= ri.size {
		err = ErrAtEnd
		return
	}

	item = ri.pointer
	ri.pointer++
	return
}

func (ri RangeIterator) At(index int) int {
	if index >= ri.size || index < 0 {
		panic(ErrNotInBounds)
	}
	return index
}

func (ri RangeIterator) Len() int {
	return ri.size
}

func (ri RangeIterator) Current() int {
	return ri.pointer
}

func (ri *RangeIterator) Reset() {
	ri.pointer = 0
}
