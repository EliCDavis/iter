package iter

func ReadFull[T any](i Iterator[T]) []T {
	vals := make([]T, 0)
	for {
		val, err := i.Next()
		if err != nil {
			break
		}
		vals = append(vals, val)
	}
	return vals
}
