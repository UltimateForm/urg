package internal

func FilterList[T comparable](source *[]T, predicate func(T) bool) ([]T, int) {
	var output []T
	var newCount int
	for _, value := range *source {
		if !predicate(value) {
			continue
		}
		output = append(output, value)
		newCount++
	}
	return output, newCount
}
