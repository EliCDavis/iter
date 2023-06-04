package iterops

import (
	"fmt"

	"github.com/EliCDavis/iter"
)

type ChunkIterator[T any] struct {
	iterator iter.Iterator[T]
	count    int
	err      error
}

func (ci *ChunkIterator[T]) Next() ([]T, error) {
	if ci.err != nil {
		return nil, ci.err
	}

	chunk := make([]T, 0, ci.count)
	var val T
	for i := 0; i < ci.count; i++ {
		val, ci.err = ci.iterator.Next()
		if ci.err != nil {
			break
		}

		chunk = append(chunk, val)
	}

	if len(chunk) > 0 {
		return chunk, nil
	}

	return nil, ci.err
}

func Chunk[T any](i iter.Iterator[T], count int) iter.Iterator[[]T] {
	if count < 1 {
		panic(fmt.Errorf("invalid chunk count: %d", count))
	}
	return &ChunkIterator[T]{
		iterator: i,
		count:    count,
	}
}
