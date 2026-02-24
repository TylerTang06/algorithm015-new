package week03

import "sort"

func permuteUnique(nums []int) [][]int {
	if nums == nil || len(nums) <= 1 {
		return [][]int{nums}
	}

	var res [][]int
	sort.Ints(nums)
	permuteUniqueRec(nums, []int{}, make([]bool, len(nums)), &res)
	return res
}

func permuteUniqueRec(nums, path []int, used []bool, res *[][]int) {
	if len(path) == len(nums) {
		*res = append(*res, path)
		return
	}

	for i := 0; i < len(nums); i++ {
		if used[i] {
			continue
		}
		if i > 0 && nums[i] == nums[i-1] && used[i-1] {
			// used[i] = true
			continue
		}

		path = append(path, nums[i])
		used[i] = true
		permuteUniqueRec(nums, path, used, res)
		used[i] = false
		path = append([]int{}, path[:len(path)-1]...)
	}

	return
}
