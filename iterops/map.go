package iterops

import "github.com/EliCDavis/iter"

type MapIterator[T, G any] struct {
	iterator       iter.Iterator[T]
	transformation func(T) G
}

func (ti MapIterator[T, G]) Next() (val G, err error) {
	var underlying T
	underlying, err = ti.iterator.Next()
	if err != nil {
		return
	}
	val = ti.transformation(underlying)
	return
}

func Map[T, G any](i iter.Iterator[T], transform func(T) G) iter.Iterator[G] {
	return MapIterator[T, G]{
		iterator:       i,
		transformation: transform,
	}
}
