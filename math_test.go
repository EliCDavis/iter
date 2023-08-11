package iter_test

import (
	"testing"

	"github.com/EliCDavis/iter"
	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	assert.Equal(t, 10, iter.Sum(iter.Range(5)))
}

func TestAverage(t *testing.T) {
	assert.Equal(t, 2., iter.Average(iter.Range(5)))
}

func TestMin(t *testing.T) {
	assert.Equal(t, 3, iter.Min(iter.Array([]int{5, 4, 3, 4, 5})))
	assert.Equal(t, -5, iter.Min(iter.Array([]int{-5, -4, -3, -4, -5})))

	assert.Equal(t, int32(3), iter.Min(iter.Array([]int32{5, 4, 3, 4, 5})))
	assert.Equal(t, int32(-5), iter.Min(iter.Array([]int32{-5, -4, -3, -4, -5})))

	assert.Equal(t, int64(3), iter.Min(iter.Array([]int64{5, 4, 3, 4, 5})))
	assert.Equal(t, int64(-5), iter.Min(iter.Array([]int64{-5, -4, -3, -4, -5})))

	assert.Equal(t, float32(3), iter.Min(iter.Array([]float32{5, 4, 3, 4, 5})))
	assert.Equal(t, float32(-5), iter.Min(iter.Array([]float32{-5, -4, -3, -4, -5})))

	assert.Equal(t, 3., iter.Min(iter.Array([]float64{5, 4, 3, 4, 5})))
	assert.Equal(t, -5., iter.Min(iter.Array([]float64{-5, -4, -3, -4, -5})))
}

func TestMax(t *testing.T) {
	assert.Equal(t, 5, iter.Max(iter.Array([]int{3, 4, 5, 4, 3})))
	assert.Equal(t, -3, iter.Max(iter.Array([]int{-3, -4, -5, -4, -3})))

	assert.Equal(t, int32(5), iter.Max(iter.Array([]int32{3, 4, 5, 4, 3})))
	assert.Equal(t, int32(-3), iter.Max(iter.Array([]int32{-3, -4, -5, -4, -3})))

	assert.Equal(t, int64(5), iter.Max(iter.Array([]int64{3, 4, 5, 4, 3})))
	assert.Equal(t, int64(-3), iter.Max(iter.Array([]int64{-3, -4, -5, -4, -3})))

	assert.Equal(t, float32(5), iter.Max(iter.Array([]float32{3, 4, 5, 4, 3})))
	assert.Equal(t, float32(-3), iter.Max(iter.Array([]float32{-3, -4, -5, -4, -3})))

	assert.Equal(t, 5., iter.Max(iter.Array([]float64{3, 4, 5, 4, 3})))
	assert.Equal(t, -3., iter.Max(iter.Array([]float64{-3, -4, -5, -4, -3})))
}
