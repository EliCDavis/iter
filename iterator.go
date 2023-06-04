package iter

type Iterator[T any] interface {
	Next() (T, error)
}

type BoundIterator[T any] interface {
	Iterator[T]
	At(index int) T
	Len() int
	Reset()
}
