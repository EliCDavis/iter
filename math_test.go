package iter_test

import (
	"testing"

	"github.com/EliCDavis/iter"
	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	assert.Equal(t, 10, iter.Sum[int](iter.Range(5)))
}

func TestAverage(t *testing.T) {
	assert.Equal(t, 2., iter.Average[int](iter.Range(5)))
}

func TestMin(t *testing.T) {
	assert.Equal(t, 3, iter.Min[int](iter.Array[int]([]int{5, 4, 3, 4, 5})))
	assert.Equal(t, -5, iter.Min[int](iter.Array[int]([]int{-5, -4, -3, -4, -5})))

	assert.Equal(t, int32(3), iter.Min[int32](iter.Array[int32]([]int32{5, 4, 3, 4, 5})))
	assert.Equal(t, int32(-5), iter.Min[int32](iter.Array[int32]([]int32{-5, -4, -3, -4, -5})))

	assert.Equal(t, int64(3), iter.Min[int64](iter.Array[int64]([]int64{5, 4, 3, 4, 5})))
	assert.Equal(t, int64(-5), iter.Min[int64](iter.Array[int64]([]int64{-5, -4, -3, -4, -5})))

	assert.Equal(t, float32(3), iter.Min[float32](iter.Array[float32]([]float32{5, 4, 3, 4, 5})))
	assert.Equal(t, float32(-5), iter.Min[float32](iter.Array[float32]([]float32{-5, -4, -3, -4, -5})))

	assert.Equal(t, 3., iter.Min[float64](iter.Array[float64]([]float64{5, 4, 3, 4, 5})))
	assert.Equal(t, -5., iter.Min[float64](iter.Array[float64]([]float64{-5, -4, -3, -4, -5})))
}

func TestMax(t *testing.T) {
	assert.Equal(t, 5, iter.Max[int](iter.Array[int]([]int{3, 4, 5, 4, 3})))
	assert.Equal(t, -3, iter.Max[int](iter.Array[int]([]int{-3, -4, -5, -4, -3})))

	assert.Equal(t, int32(5), iter.Max[int32](iter.Array[int32]([]int32{3, 4, 5, 4, 3})))
	assert.Equal(t, int32(-3), iter.Max[int32](iter.Array[int32]([]int32{-3, -4, -5, -4, -3})))

	assert.Equal(t, int64(5), iter.Max[int64](iter.Array[int64]([]int64{3, 4, 5, 4, 3})))
	assert.Equal(t, int64(-3), iter.Max[int64](iter.Array[int64]([]int64{-3, -4, -5, -4, -3})))

	assert.Equal(t, float32(5), iter.Max[float32](iter.Array[float32]([]float32{3, 4, 5, 4, 3})))
	assert.Equal(t, float32(-3), iter.Max[float32](iter.Array[float32]([]float32{-3, -4, -5, -4, -3})))

	assert.Equal(t, 5., iter.Max[float64](iter.Array[float64]([]float64{3, 4, 5, 4, 3})))
	assert.Equal(t, -3., iter.Max[float64](iter.Array[float64]([]float64{-3, -4, -5, -4, -3})))
}
