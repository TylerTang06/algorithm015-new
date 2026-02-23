package week01

/*
189.轮转数组
给定一个整数数组 nums，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。

示例 1:

输入: nums = [1,2,3,4,5,6,7], k = 3
输出: [5,6,7,1,2,3,4]
解释:
向右轮转 1 步: [7,1,2,3,4,5,6]
向右轮转 2 步: [6,7,1,2,3,4,5]
向右轮转 3 步: [5,6,7,1,2,3,4]
示例 2:

输入：nums = [-1,-100,3,99], k = 2
输出：[3,99,-1,-100]
解释:
向右轮转 1 步: [99,-1,-100,3]
向右轮转 2 步: [3,99,-1,-100]


提示：

1 <= nums.length <= 10^5
-231 <= nums[i] <= 231 - 1
0 <= k <= 10^5
*/

func rotate(nums []int, k int) {
	if len(nums) <= 1 || k == 0 || k == len(nums) {
		return
	}

	count := 0
	for stIndex := 0; count < len(nums); stIndex++ {
		temp, nxtIndx := nums[stIndex], (k+stIndex)%len(nums)
		for stIndex != nxtIndx {
			nums[nxtIndx], temp = temp, nums[nxtIndx]
			nxtIndx = (nxtIndx + k) % len(nums)
			count++
		}
		nums[nxtIndx] = temp
		count++
	}
}
