package iterops_test

import (
	"strconv"
	"testing"

	"github.com/EliCDavis/iter"
	"github.com/EliCDavis/iter/iterops"
	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	originalData := []string{"1", "2", "3", "4", "5"}
	i := iter.Array(originalData)

	mapper := iterops.Map[string, int](&i, func(s string) int {
		parsed, _ := strconv.Atoi(s)
		return parsed
	})

	mappedData := iter.ReadFull(mapper)

	assert.Equal(t, []int{1, 2, 3, 4, 5}, mappedData)
}
