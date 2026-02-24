package week08

import "sort"

func relativeSortArray(arr1 []int, arr2 []int) []int {
	if arr2 == nil || len(arr2) == 0 {
		sort.Ints(arr1)
		return arr1
	}

	myMap := make(map[int]int)
	for _, num := range arr1 {
		if _, ok := myMap[num]; ok {
			myMap[num]++
		} else {
			myMap[num] = 1
		}
	}

	i := 0
	for _, num := range arr2 {
		count := myMap[num]
		delete(myMap, num)
		for count > 0 {
			arr1[i] = num
			i++
			count--
		}
	}

	j := i
	for num, count := range myMap {
		for count > 0 {
			arr1[j] = num
			j++
			count--
		}
	}

	sort.Ints(arr1[i:])

	return arr1
}
