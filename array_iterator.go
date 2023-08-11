package iter

type ArrayIterator[T any] struct {
	data    []T
	pointer int
}

func Array[T any](data []T) *ArrayIterator[T] {
	return &ArrayIterator[T]{
		data: data,
	}
}

func (i *ArrayIterator[T]) Next() (item T, err error) {
	if i.pointer >= len(i.data) {
		err = ErrAtEnd
		return
	}

	item = i.data[i.pointer]
	i.pointer++
	return
}

func (i ArrayIterator[T]) At(index int) T {
	return i.data[index]
}

func (i ArrayIterator[T]) Len() int {
	return len(i.data)
}

func (i ArrayIterator[T]) Current() T {
	return i.data[i.pointer]
}

func (i *ArrayIterator[T]) Reset() {
	i.pointer = 0
}
