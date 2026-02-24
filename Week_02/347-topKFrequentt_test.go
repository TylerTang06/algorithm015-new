package week02

import (
	"fmt"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	// [-1, -1] 1
	res := topKFrequent([]int{-1, -1}, 1)
	fmt.Println("1. ", res)
	res = []int{}
	// [-1, 1, 4, -4, 3, 5, 4, -2, 3, -1] 3
	res = topKFrequent([]int{-1, 1, 4, -4, 3, 5, 4, -2, 3, -1}, 3)
	fmt.Println("2. ", res)
}
