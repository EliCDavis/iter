package iterops_test

import (
	"fmt"
	"testing"

	"github.com/EliCDavis/iter"
	"github.com/EliCDavis/iter/iterops"
	"github.com/stretchr/testify/assert"
)

func TestZip(t *testing.T) {
	arrA := iter.Array([]int{1, 2, 3, 4, 5})
	arrB := iter.Array([]string{"A", "B", "C", "D", "E"})

	mapper := iterops.Zip[int, string, string](&arrA, &arrB, func(i int, s string) string {
		return fmt.Sprintf("%d%s", i, s)
	})

	mappedData := iter.ReadFull[string](mapper)

	assert.Equal(t, []string{"1A", "2B", "3C", "4D", "5E"}, mappedData)
}

func TestZipStreamEndsEarly(t *testing.T) {
	arrA := iter.Array([]int{1, 2, 3, 4, 5})
	arrB := iter.Array([]string{"A", "B", "C", "D"})

	mapper := iterops.Zip[int, string, string](&arrA, &arrB, func(i int, s string) string {
		return fmt.Sprintf("%d%s", i, s)
	})

	mappedData := iter.ReadFull[string](mapper)

	assert.Equal(t, []string{"1A", "2B", "3C", "4D"}, mappedData)
}
