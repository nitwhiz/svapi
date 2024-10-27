package util

import "strings"

func Combinations(values ...string) []string {
	var result []string

	n := len(values)
	totalCombinations := 1 << n

	for i := 0; i < totalCombinations; i++ {
		var combination []string
		hasNonEmpty := false

		for j := 0; j < n; j++ {
			if i&(1<<j) != 0 {
				combination = append(combination, values[j])
				hasNonEmpty = true
			} else {
				combination = append(combination, "")
			}
		}

		if hasNonEmpty {
			result = append(result, strings.Join(combination, "_"))
		}
	}

	return result
}
