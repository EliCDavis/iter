package iter_test

import (
	"testing"

	"github.com/EliCDavis/iter"
	"github.com/stretchr/testify/assert"
)

func TestReadFull(t *testing.T) {
	original := []int{1, 2, 3, 4, 5}
	i := iter.Array(original)
	assert.Equal(t, original, iter.ReadFull[int](i))
}
