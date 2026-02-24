package week03

/*
46.全排列
给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。

示例 1：

输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
示例 2：

输入：nums = [0,1]
输出：[[0,1],[1,0]]
示例 3：

输入：nums = [1]
输出：[[1]]


提示：

1 <= nums.length <= 6
-10 <= nums[i] <= 10
nums 中的所有整数 互不相同
*/

func permute(nums []int) [][]int {
	if len(nums) <= 1 {
		return [][]int{nums}
	}

	used := make([]bool, len(nums))
	for i := 0; i < len(used); i++ {
		used[i] = false
	}

	return recursion(nums, []int{}, used, [][]int{})
}

func recursion(nums, path []int, used []bool, res [][]int) [][]int {
	if len(nums) == len(path) {
		res = append(res, path)
		return res
	}

	for i := 0; i < len(nums); i++ {
		if !used[i] {
			used[i] = true
			path = append(path, nums[i])
			res = recursion(nums, path, used, res)
			used[i] = false
			// path = path[:len(path)-1] not right
			path = append([]int{}, path[:len(path)-1]...)
		}
	}

	return res
}

// func permute(nums []int) [][]int {
// 	if len(nums) <= 1 {
// 		return [][]int{nums}
// 	}

// 	n := len(nums)
// 	used := make([]bool, n)

// 	var res [][]int
// 	var recursion func(path []int)
// 	recursion = func(path []int) {
// 		if len(path) == n {
// 			res = append(res, path)
// 			return
// 		}

// 		for i := 0; i < n; i++ {
// 			if !used[i] {
// 				used[i] = true
// 				path = append(path, nums[i])
// 				recursion(path)
// 				used[i] = false
// 				path = append([]int{}, path[:len(path)-1]...)
// 			}
// 		}
// 	}

// 	recursion([]int{})
// 	return res
// }
