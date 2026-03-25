package week04

// 33.搜索旋转排序数组
/*
整数数组 nums 按升序排列，数组中的值 互不相同 。

在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了 向左旋转，使数组变为 [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始 计数）。例如， [0,1,2,4,5,6,7] 下标 3 上向左旋转后可能变为 [4,5,6,7,0,1,2] 。

给你 旋转后 的数组 nums 和一个整数 target ，如果 nums 中存在这个目标值 target ，则返回它的下标，否则返回 -1 。

你必须设计一个时间复杂度为 O(log n) 的算法解决此问题。

示例 1：

输入：nums = [4,5,6,7,0,1,2], target = 0
输出：4
示例 2：

输入：nums = [4,5,6,7,0,1,2], target = 3
输出：-1
示例 3：

输入：nums = [1], target = 0
输出：-1
*/
func search(nums []int, target int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}

	l, r := 0, len(nums)-1
	for l < r {
		mid := l + (r-l)/2
		// ok
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < nums[r] {
			// nums[mid]<target<nums[r]
			if target > nums[mid] && target <= nums[r] {
				l = mid + 1
			} else {
				// nums[l]->target->nums[mid]<nums[r]
				r = mid - 1
			}
		} else {
			// nums[l]<target<nums[mid]
			if target >= nums[l] && target < nums[mid] {
				r = mid - 1
			} else {
				// nums[l]->nums[mid]->target->nums[r]
				l = mid + 1
			}
		}
	}

	if nums[l] == target {
		return l
	}

	return -1
}
