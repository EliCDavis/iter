package iter

type ArrayIterator[T any] struct {
	data    []T
	pointer int
}

func Array[T any](data []T) ArrayIterator[T] {
	return ArrayIterator[T]{
		data: data,
	}
}

func (i ArrayIterator[T]) Continue() bool {
	return i.pointer < len(i.data)
}

func (i *ArrayIterator[T]) Next() T {
	if !i.Continue() {
		panic(ErrAtEnd)
	}
	item := i.data[i.pointer]
	i.pointer++
	return item
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
