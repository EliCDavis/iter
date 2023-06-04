package iterops_test

import (
	"testing"

	"github.com/EliCDavis/iter"
	"github.com/EliCDavis/iter/iterops"
	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	arr := iter.Array([]int{1, 2, 3, 4, 5})

	mapper := iterops.Filter[int](&arr, func(i int) bool {
		return i % 2 == 0
	})

	mappedData := iter.ReadFull(mapper)

	assert.Equal(t, []int{2, 4}, mappedData)
}
