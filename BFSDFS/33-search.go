package week04

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
