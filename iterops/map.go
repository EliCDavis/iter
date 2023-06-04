package iterops

import "github.com/EliCDavis/iter"

type MapIterator[T, G any] struct {
	iterator       iter.Iterator[T]
	transformation func(T) G
}

func (ti MapIterator[T, G]) Next() G {
	return ti.transformation(ti.iterator.Next())
}

func (ti MapIterator[T, G]) Continue() bool {
	return ti.iterator.Continue()
}

func Map[T, G any](i iter.Iterator[T], transform func(T) G) iter.Iterator[G] {
	return MapIterator[T, G]{
		iterator:       i,
		transformation: transform,
	}
}
