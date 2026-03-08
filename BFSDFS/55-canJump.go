package week04

func canJump(nums []int) bool {
	if nums == nil || len(nums) == 0 {
		return false
	}

	// greedy, it is good idea
	maxStep := 0
	for i := 0; i < len(nums); i++ {
		if i > maxStep {
			return false
		}
		if maxStep < nums[i]+i {
			maxStep = nums[i] + i
		}
	}

	return true
}
