package week02

import (
	"container/heap"
)

/*
374.前K个高频词汇
给你一个整数数组 nums 和一个整数 k ，请你返回其中出现频率前 k 高的元素。你可以按 任意顺序 返回答案。



示例 1：

输入：nums = [1,1,1,2,2,3], k = 2

输出：[1,2]

示例 2：

输入：nums = [1], k = 1

输出：[1]

示例 3：

输入：nums = [1,2,1,2,1,2,3,1,3,2], k = 2

输出：[1,2]
*/

func topKFrequent(nums []int, k int) []int {
	// 1. 统计频率
	freqMap := make(map[int]int)
	for _, num := range nums {
		freqMap[num]++
	}

	// 2. 创建小顶堆（按频率排序）
	h := &IntHeap{}
	heap.Init(h)

	// 3. 维护大小为k的堆
	for num, freq := range freqMap {
		heap.Push(h, [2]int{num, freq})
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	// 4. 收集结果
	result := make([]int, k)
	for i := k - 1; i >= 0; i-- {
		result[i] = heap.Pop(h).([2]int)[0]
	}

	return result
}

// 小顶堆，存储 [num, freq]
type IntHeap [][2]int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i][1] < h[j][1] } // 按频率升序
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1] // 移除最后一个元素
	*h = old[0 : n-1]
	return x
}
