package iter

type Iterator[T any] interface {
	Continue() bool
	Next() T
}

type BoundIterator[T any] interface {
	Iterator[T]
	At(index int) T
	Len() int
	Reset()
}
