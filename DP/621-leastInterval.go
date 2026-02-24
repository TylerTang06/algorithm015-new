package week06

import "sort"

func leastInterval(tasks []byte, n int) int {
	if tasks == nil || len(tasks) == 0 {
		return 0
	}

	arr := make([]int, 26)
	for i := 0; i < len(tasks); i++ {
		arr[tasks[i]-'A']++
	}
	sort.Ints(arr)

	max := arr[25] - 1
	allSlots := max * n
	for i := 24; i >= 0 && arr[i] > 0; i-- {
		if arr[i] < max {
			allSlots -= arr[i]
		} else {
			allSlots -= max
		}
	}

	if allSlots > 0 {
		return allSlots + len(tasks)
	}
	return len(tasks)
}
