package iterops_test

import (
	"testing"

	"github.com/EliCDavis/iter"
	"github.com/EliCDavis/iter/iterops"
	"github.com/stretchr/testify/assert"
)

func TestFlatten(t *testing.T) {
	arr := iter.Array([][]int{
		{},
		{1, 2, 3, 4},
		{},
		{5, 6, 7, 8},
		{},
		{},
	})

	flatten := iterops.Flatten[int](&arr)

	flattenData := iter.ReadFull(flatten)

	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8}, flattenData)
}
