package utils

func Unique[T comparable](input []T) []T {
	seen := make(map[T]struct{})
	result := make([]T, 0)

	for _, v := range input {
		if _, exists := seen[v]; !exists {
			seen[v] = struct{}{}
			result = append(result, v)
		}
	}

	return result
}
