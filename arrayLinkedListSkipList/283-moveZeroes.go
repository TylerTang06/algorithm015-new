package week01

/*
283.移动0

给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

请注意 ，必须在不复制数组的情况下原地对数组进行操作。
*/

func moveZeroes(nums []int) {
	if nums == nil || len(nums) <= 1 {
		return
	}

	for i, j := 0, 0; j < len(nums); j++ {
		if nums[j] == 0 {
			continue
		}
		nums[i], nums[j] = nums[j], nums[i]
		i++
	}
}
