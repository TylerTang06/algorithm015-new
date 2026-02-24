package week03

func combine(n int, k int) [][]int {
	res := [][]int{}
	// recursion terminator
	if k == 1 {
		for i := 0; i < n; i++ {
			res = append(res, []int{i + 1})
		}
		return res
	}

	// subproblems
	for i := n; i >= k; i-- {
		r := combine(i-1, k-1)
		// process result
		for j := 0; j < len(r); j++ {
			r[j] = append(r[j], i)
			res = append(res, r[j])
		}
	}

	return res
}
