# Iter
![Coverage](https://img.shields.io/badge/Coverage-100.0%25-brightgreen)

Iterator and utilities.

## Examples

### Iterating

Iterate over an iterator

```golang
package main

import (
    "fmt"

    "github.com/EliCDavis/iter"
)

func main() {
	arrayIter := iter.Array([]int{1, 2, 3, 4, 5})

    sum := 0
	for arrayIter.Continue() {
		sum += arrayIter.Next()
	}

    // Prints: 15
    fmt.Println(sum)
}

```

### Mapping

Map an iterator of one type to an iterator of a different type.

```golang
package main

import (
    "fmt"
	"strconv"

	"github.com/EliCDavis/iter"
	"github.com/EliCDavis/iter/iterops"
)

func main() {
	arr := iter.Array([]string{"1", "2", "3", "4", "5"})

	mapper := iterops.Map[string, int](&arr, func(s string) int {
		parsed, err := strconv.Atoi(s)
        if err != nil {
            panic(err)
        }
		return parsed
	})

    // Prints: [1 2 3 4 5]
    fmt.Println(iter.ReadFull(mapper))
}
```
