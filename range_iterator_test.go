package iter_test

import (
	"testing"

	"github.com/EliCDavis/iter"
	"github.com/stretchr/testify/assert"
)

func TestRange(t *testing.T) {
	i := iter.Range(5)

	// test Current
	assert.Equal(t, 0, i.Current())

	// test reading all elements
	assert.Equal(t, []int{0, 1, 2, 3, 4}, iter.ReadFull[int](i))

	// test at
	assert.Equal(t, 0, i.At(0))
	assert.Equal(t, 1, i.At(1))
	assert.Equal(t, 2, i.At(2))
	assert.Equal(t, 3, i.At(3))
	assert.Equal(t, 4, i.At(4))
	assert.Equal(t, 5, i.Len())

	// test reset
	assert.Equal(t, 5, i.Current())
	i.Reset()
	assert.Equal(t, 0, i.Current())

	// test upper bound At()
	assert.PanicsWithError(t, iter.ErrNotInBounds.Error(), func() {
		i.At(5)
	})

	// test lower bound At()
	assert.PanicsWithError(t, iter.ErrNotInBounds.Error(), func() {
		i.At(-1)
	})
}

func TestRangePanicsWithValueLessThan0(t *testing.T) {
	assert.PanicsWithError(t, iter.ErrInvalidRangeValue.Error(), func() {
		iter.Range(-1)
	})
}
