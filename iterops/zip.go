package iterops

import "github.com/EliCDavis/iter"

type ZipIterator[A, B, T any] struct {
	iteratorA      iter.Iterator[A]
	iteratorB      iter.Iterator[B]
	transformation func(A, B) T
}

func (ti ZipIterator[A, B, T]) Next() (val T, err error) {
	var underlyingA A
	underlyingA, err = ti.iteratorA.Next()
	if err != nil {
		return
	}

	var underlyingB B
	underlyingB, err = ti.iteratorB.Next()
	if err != nil {
		return
	}

	val = ti.transformation(underlyingA, underlyingB)
	return
}

func Zip[A, B, T any](a iter.Iterator[A], b iter.Iterator[B], transform func(A, B) T) ZipIterator[A, B, T] {
	return ZipIterator[A, B, T]{
		iteratorA:      a,
		iteratorB:      b,
		transformation: transform,
	}
}
