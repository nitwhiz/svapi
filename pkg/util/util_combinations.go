package util

import "strings"

func Combinations(values ...string) []string {
	var result []string

	seen := map[string]struct{}{}
	n := len(values)
	totalCombinations := 1 << n

	for i := 0; i < totalCombinations; i++ {
		var combination []string
		hasNonEmpty := false

		for j := 0; j < n; j++ {
			if i&(1<<j) != 0 {
				combination = append(combination, values[j])

				if values[j] != "" {
					hasNonEmpty = true
				}
			} else {
				combination = append(combination, "")
			}
		}

		if hasNonEmpty {
			combinationString := strings.Join(combination, "_")

			if _, ok := seen[combinationString]; ok {
				continue
			}

			seen[combinationString] = struct{}{}

			result = append(result, combinationString)
		}
	}

	return result
}
