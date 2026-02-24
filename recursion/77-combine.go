package week03

/*
77.组合
给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。

你可以按 任何顺序 返回答案。

示例 1：

输入：n = 4, k = 2
输出：
[
  [2,4],
  [3,4],
  [2,3],
  [1,2],
  [1,3],
  [1,4],
]
示例 2：

输入：n = 1, k = 1
输出：[[1]]


提示：

1 <= n <= 20
1 <= k <= n

*/

// 递归 会重复计算
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
		// recursion
		r := combine(i-1, k-1)
		// process result
		for j := 0; j < len(r); j++ {
			r[j] = append(r[j], i)
			res = append(res, r[j])
		}
	}

	return res
}

// 动态规划：从小问题到大问题，避免重复计算
// func combine(n int, k int) [][]int {
// 	var res [][]int
// 	var temp []int
// 	var dfs func(int)
// 	dfs = func(cur int) {
// 		// 剪枝：temp的长度加上[cur,n]的长度小于k，不可能构建疮毒为k的temp
// 		if len(temp)+(n-cur+1) < k {
// 			return
// 		}

// 		// terminator
// 		if len(temp) == k {
// 			r := make([]int, k)
// 			copy(r, temp)
// 			res = append(res, r)
// 			return
// 		}

// 		// 选择当前元素
// 		temp = append(temp, cur)
// 		dfs(cur + 1)
// 		temp = temp[:len(temp)-1]
// 		// 不考虑当前元素
// 		dfs(cur + 1)
// 	}

// 	dfs(1)

// 	return res
// }
