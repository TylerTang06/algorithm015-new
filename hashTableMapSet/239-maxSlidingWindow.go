package week02

import (
	"container/heap"
	"container/list"
	"sort"
)

/*
239.滑动窗口最大值

给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。
你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。

返回 滑动窗口中的最大值 。

示例 1：

输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
输出：[3,3,5,5,6,7]
解释：
滑动窗口的位置                最大值
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7
示例 2：

输入：nums = [1], k = 1
输出：[1]


提示：

1 <= nums.length <= 10^5
-10^4 <= nums[i] <= 10^4
1 <= k <= nums.length
*/

// 单调队列法
func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) < k {
		return []int{}
	}

	res := []int{}
	// can use arry to be queue
	myQue := list.New()
	for i := 0; i < len(nums); i++ {
		if i > k-1 && myQue.Front().Value.(int) == i-k {
			myQue.Remove(myQue.Front())
		}
		// nums[i] compare with back value of Queue
		for myQue.Len() > 0 && nums[myQue.Back().Value.(int)] <= nums[i] {
			myQue.Remove(myQue.Back())
		}

		myQue.PushBack(i)
		if i >= k-1 {
			res = append(res, nums[myQue.Front().Value.(int)])
		}
	}

	return res
}

// 使用大根堆
var hpSi []int

type hp struct {
	sort.IntSlice
}

func (h hp) Less(i, j int) bool {
	return hpSi[h.IntSlice[i]] > hpSi[h.IntSlice[j]]
}

func (h *hp) Push(v interface{}) {
	h.IntSlice = append(h.IntSlice, v.(int))
}

func (h *hp) Pop() interface{} {
	s := h.IntSlice
	v := s[len(s)-1]
	h.IntSlice = s[:len(s)-1]
	return v
}

func maxSlidingWindow1(nums []int, k int) []int {
	hpSi = nums
	q := &hp{make([]int, k)}

	for i := 0; i < k; i++ {
		q.IntSlice[i] = i
	}
	heap.Init(q)

	var res []int
	res = append(res, nums[q.IntSlice[0]])
	for i := k; i < len(nums); i++ {
		heap.Push(q, i)
		for q.IntSlice[0] <= i-k {
			heap.Pop(q)
		}
		res = append(res, nums[q.IntSlice[0]])
	}

	return res
}
