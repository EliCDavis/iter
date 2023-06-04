package iter_test

import (
	"testing"

	"github.com/EliCDavis/iter"
	"github.com/stretchr/testify/assert"
)

func TestArrayIterator(t *testing.T) {
	originalData := []int{1, 2, 3, 4, 5}
	arrayIter := iter.Array(originalData)

	assert.Equal(t, originalData[0], arrayIter.At(0))
	assert.Equal(t, originalData[1], arrayIter.At(1))
	assert.Equal(t, originalData[2], arrayIter.At(2))
	assert.Equal(t, originalData[3], arrayIter.At(3))
	assert.Equal(t, originalData[4], arrayIter.At(4))
	assert.Equal(t, 5, arrayIter.Len())
	assert.Equal(t, 1, arrayIter.Current())

	newData := make([]int, 0, len(originalData))
	for {
		item, err := arrayIter.Next()
		if err != nil {
			assert.EqualError(t, err, iter.ErrAtEnd.Error())
			break
		}
		newData = append(newData, item)
	}
	assert.Equal(t, originalData, newData)

	arrayIter.Reset()
	assert.Equal(t, 1, arrayIter.Current())
}
