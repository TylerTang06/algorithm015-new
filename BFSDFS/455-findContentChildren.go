package week04

import "sort"

func findContentChildren(g []int, s []int) int {
	if g == nil || s == nil || len(g) == 0 || len(s) == 0 {
		return 0
	}

	// greedy, sort fistly
	sort.Ints(g)
	sort.Ints(s)
	i, j := 0, 0
	for i < len(g) && j < len(s) {
		if g[i] <= s[j] {
			i++
			j++
		} else {
			j++
		}
	}

	return i
}
