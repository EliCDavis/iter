# Iter

![Coverage](https://img.shields.io/badge/Coverage-100.0%25-brightgreen)
[![Go Report Card](https://goreportcard.com/badge/github.com/EliCDavis/iter)](https://goreportcard.com/report/github.com/EliCDavis/iter)

Iterator and utilities.

Some inspiration from [ReactiveX](https://reactivex.io/).

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
    arr := iter.Array([]int{1, 2, 3, 4, 5})

    sum := 0
    for {
        val, err := arr.Next()
        if err != nil {
            break
        }
        sum += val
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

### Range

Range is a bound iteractor that counts up from 0 to n-1 the value specified

```golang
package main

import (
    "fmt"

    "github.com/EliCDavis/iter"
)

func main() {
    arr := iter.Range(5)

    sum := 0
    for {
        val, err := arr.Next()
        if err != nil {
            break
        }
        sum += val
    }

    // Prints: 10 (0 + 1 + 2 + 3 + 4)
    fmt.Println(sum)
}
```