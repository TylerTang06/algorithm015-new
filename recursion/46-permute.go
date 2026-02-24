package week03

func permute(nums []int) [][]int {
	if nums == nil || len(nums) <= 1 {
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
