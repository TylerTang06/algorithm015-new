package week04

func findMin(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		return nums[0]
	}

	l, r := 0, len(nums)-1
	// when l==r, it's break and its the min point
	for l < r {
		mid := l + (r-l)/2
		// [2,1] or [1,2]
		if nums[mid] == nums[l] || nums[mid] == nums[r] {
			if nums[l] < nums[r] {
				return nums[l]
			}
			return nums[r]
		}
		if nums[mid] > nums[l] {
			if nums[mid] > nums[r] {
				l = mid
			} else if nums[mid] < nums[r] {
				r = mid
			}
		}
		if nums[mid] < nums[l] {
			if nums[mid] < nums[r] {
				r = mid
			}
		}
	}

	return nums[l]
}
