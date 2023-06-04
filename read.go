package iter

func ReadFull[T any](i Iterator[T]) []T {
	vals := make([]T, 0)
	for i.Continue() {
		vals = append(vals, i.Next())
	}
	return vals
}
