package iterops_test

import (
	"testing"

	"github.com/EliCDavis/iter"
	"github.com/EliCDavis/iter/iterops"
	"github.com/stretchr/testify/assert"
)

func TestScan(t *testing.T) {
	arr := iter.Array([]int{1, 2, 3, 4})

	mapper := iterops.Scan[int](arr, func(cur, acc int) int {
		return cur + acc
	})

	mappedData := iter.ReadFull[int](&mapper)

	assert.Equal(t, []int{1, 3, 6, 10}, mappedData)
}
