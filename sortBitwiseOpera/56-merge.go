package week08

import "sort"

func merge(intervals [][]int) [][]int {
	if intervals == nil || len(intervals) == 0 || len(intervals[0]) <= 1 {
		return [][]int{}
	}

	sort.Sort(sortBy(intervals))
	res := [][]int{}
	index := -1
	for _, interval := range intervals {
		if index == -1 || res[index][1] < interval[0] {
			res = append(res, interval)
			index++
		} else if res[index][1] < interval[1] {
			res[index][1] = interval[1]
		}
	}

	return res
}

type sortBy [][]int

func (a sortBy) Len() int           { return len(a) }
func (a sortBy) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sortBy) Less(i, j int) bool { return a[i][0] < a[j][0] }
