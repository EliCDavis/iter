package iterops_test

import (
	"testing"

	"github.com/EliCDavis/iter"
	"github.com/EliCDavis/iter/iterops"
	"github.com/stretchr/testify/assert"
)

func TestChunk(t *testing.T) {
	arr := iter.Array([]int{1, 2, 3, 4, 5})
	chunkedIterator := iterops.Chunk[int](&arr, 2)

	mappedData := iter.ReadFull(chunkedIterator)

	assert.Len(t, mappedData, 3)
	assert.Equal(t, []int{1, 2}, mappedData[0])
	assert.Equal(t, []int{3, 4}, mappedData[1])
	assert.Equal(t, []int{5}, mappedData[2])
}

func TestChunkPerfect(t *testing.T) {
	arr := iter.Array([]int{1, 2, 3, 4, 5, 6})
	chunkedIterator := iterops.Chunk[int](&arr, 2)

	mappedData := iter.ReadFull(chunkedIterator)

	assert.Len(t, mappedData, 3)
	assert.Equal(t, []int{1, 2}, mappedData[0])
	assert.Equal(t, []int{3, 4}, mappedData[1])
	assert.Equal(t, []int{5, 6}, mappedData[2])
}

func TestChunkThrowsWithCountLessThanOne(t *testing.T) {
	assert.PanicsWithError(t, "invalid chunk count: 0", func() {
		iterops.Chunk[int](nil, 0)
	})
}
