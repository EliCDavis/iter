package iter

import "math"

type Number interface {
	float32 | float64 | int | int32 | int64
}

func Sum[T Number](i Iterator[T]) T {
	var sum T
	for {
		val, err := i.Next()
		if err != nil {
			break
		}
		sum += val
	}
	return sum
}

func Min[T Number](i Iterator[T]) T {
	var min T

	switch i.(type) {
	case Iterator[int]:
		max := math.MaxInt
		min = T(max)

	case Iterator[int32]:
		min = T(math.MaxInt32)

	case Iterator[int64]:
		max := math.MaxInt64
		min = T(max)

	case Iterator[float32]:
		max := math.MaxFloat32
		min = T(max)

	case Iterator[float64]:
		max := math.MaxFloat64
		min = T(max)
	}

	for {
		val, err := i.Next()
		if err != nil {
			break
		}
		if min > val {
			min = val
		}
	}
	return min
}

func Max[T Number](i Iterator[T]) T {
	var max T

	switch i.(type) {
	case Iterator[int]:
		min := math.MinInt
		max = T(min)

	case Iterator[int32]:
		max = T(math.MinInt32)

	case Iterator[int64]:
		min := math.MinInt64
		max = T(min)

	case Iterator[float32]:
		min := -math.MaxFloat32
		max = T(min)

	case Iterator[float64]:
		min := -math.MaxFloat64
		max = T(min)
	}

	for {
		val, err := i.Next()
		if err != nil {
			break
		}
		if max < val {
			max = val
		}

	}
	return max
}

func Average[T Number](i Iterator[T]) float64 {
	var sum T
	count := 0.
	for {
		val, err := i.Next()
		if err != nil {
			break
		}
		sum += val
		count++
	}
	return float64(sum) / count
}
