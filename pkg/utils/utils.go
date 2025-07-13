package utils

import "strings"

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

func StringToBool(s string) bool {
	lower := strings.ToLower(s)
	if lower == "true" {
		return true
	}
	return false
}
