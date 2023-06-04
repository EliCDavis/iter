package iterops

import "github.com/EliCDavis/iter"

type FlattenIterator[T any] struct {
	iterator iter.Iterator[[]T]
	cur      []T
	i        int
}

func (fi *FlattenIterator[T]) Next() (val T, err error) {
	if fi.i < len(fi.cur) {
		val = fi.cur[fi.i]
		fi.i++
		return
	}

	fi.i = 0
	for {
		fi.cur, err = fi.iterator.Next()
		if err != nil {
			return
		}

		if len(fi.cur) > 0 {
			break
		}
	}

	val = fi.cur[fi.i]
	fi.i++
	return
}

func Flatten[T any](i iter.Iterator[[]T]) iter.Iterator[T] {
	return &FlattenIterator[T]{
		iterator: i,
	}
}
