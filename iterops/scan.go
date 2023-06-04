package iterops

import "github.com/EliCDavis/iter"

type ScanIterator[T any] struct {
	iterator       iter.Iterator[T]
	transformation func(T, T) T
	acc            T
}

func (ti *ScanIterator[T]) Next() (val T, err error) {
	var next T
	next, err = ti.iterator.Next()
	if err != nil {
		return
	}
	ti.acc = ti.transformation(next, ti.acc)
	return ti.acc, nil
}

func Scan[T any](i iter.Iterator[T], transform func(T, T) T) ScanIterator[T] {
	return ScanIterator[T]{
		iterator:       i,
		transformation: transform,
	}
}
