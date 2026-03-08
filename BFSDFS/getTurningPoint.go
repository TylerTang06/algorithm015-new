package week04

func getTurningPoint(nums []int) int {
	if nums == nil || len(nums) <= 1 {
		return 0
	}
	if nums[0] < nums[len(nums)-1] {
		return 0
	}

	l, r := 0, len(nums)-1
	for l < r {
		mid := l + (r-l)/2
		if nums[mid] == nums[l] || nums[mid] == nums[r] {
			if nums[l] > nums[r] {
				return nums[r]
			}
			return nums[l]
		}
		if nums[mid] > nums[l] {
			if nums[mid] < nums[r] {
				return nums[l]
			}
			l = mid + 1
		} else {
			if nums[mid] < nums[r] {
				r = mid - 1
			}
		}
	}

	return nums[l]
}
