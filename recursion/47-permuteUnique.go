package week03

import "sort"

/*
47.全排列2
给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。

示例 1：

输入：nums = [1,1,2]
输出：
[[1,1,2],
 [1,2,1],
 [2,1,1]]
示例 2：

输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]


提示：

1 <= nums.length <= 8
-10 <= nums[i] <= 10
*/
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

// func permuteUnique(nums []int) [][]int {
// 	if len(nums) <= 1 {
// 		return [][]int{nums}
// 	}

// 	n := len(nums)
// 	used := make([]bool, n)
// 	sort.Ints(nums)

// 	var res [][]int
// 	var recursion func(path []int)
// 	recursion = func(path []int) {
// 		if len(path) == n {
// 			res = append(res, path)
// 			return
// 		}

// 		for i := 0; i < n; i++ {
// 			if used[i] || (i > 0 && nums[i] == nums[i-1] && used[i-1]) {
// 				continue
// 			}

// 			used[i] = true
// 			path = append(path, nums[i])
// 			recursion(path)
// 			used[i] = false
// 			path = append([]int{}, path[:len(path)-1]...)
// 		}
// 	}

// 	recursion([]int{})
// 	return res
// }
