package internal

func FilterList[T comparable](source *[]T, predicate func(T) bool) ([]T, int) {
	var output []T
	for _, value := range *source {
		if !predicate(value) {
			continue
		}
		output = append(output, value)
	}
	return output, len(output)
}
