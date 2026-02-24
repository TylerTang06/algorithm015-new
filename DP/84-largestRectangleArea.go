package week06

import "container/list"

func largestRectangleArea(heights []int) int {
	if heights == nil || len(heights) == 0 {
		return 0
	}

	res := 0
	heights = append([]int{0}, heights...)
	heights = append(heights, 0)
	stk := list.New()
	stk.PushBack(0)
	for i := 1; i < len(heights); i++ {
		for heights[stk.Back().Value.(int)] > heights[i] {
			h := heights[stk.Back().Value.(int)]
			stk.Remove(stk.Back())
			w := i - stk.Back().Value.(int) - 1
			if h*w > res {
				res = h * w
			}
		}
		stk.PushBack(i)
	}

	return res
}
