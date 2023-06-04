package iterops

import "github.com/EliCDavis/iter"

type FilterIterator[T any] struct {
	iterator iter.Iterator[T]
	filter   func(T) bool
}

func (fi FilterIterator[T]) Next() (val T, err error) {
	for {
		val, err = fi.iterator.Next()
		if err != nil {
			return
		}

		if fi.filter(val) {
			return
		}
	}
}

func Filter[T any](i iter.Iterator[T], filter func(T) bool) iter.Iterator[T] {
	return FilterIterator[T]{
		iterator: i,
		filter:   filter,
	}
}
