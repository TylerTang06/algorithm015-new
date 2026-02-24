package week08

func solveNQueens(n int) [][]string {
	if n < 1 {
		return [][]string{}
	}

	res := solveNQueensUseBits(n, 0, 0, 0, 0, []int{}, [][]int{})

	return generateNQueensResults(n, res)
}

func solveNQueensUseBits(n, row, cols, sum, diff int, cur []int, res [][]int) [][]int {
	if row >= n {
		res = append(res, cur)
		return res
	}

	bits := (^(cols | sum | diff)) & ((1 << n) - 1) // 得到当前所有空位
	for bits != 0 {
		p := bits & -bits // 取到最低位 1
		index, tmp := 0, p
		for tmp > 0 {
			index++
			tmp >>= 1
		}
		cur = append(cur, index)
		res = solveNQueensUseBits(n, row+1, cols|p, (sum|p)<<1, (diff|p)>>1, cur, res)
		bits = bits & (bits - 1) // 去掉最低位 1
		cur = append([]int{}, cur[:len(cur)-1]...)
	}

	return res
}

func generateNQueensResults(n int, input [][]int) (results [][]string) {
	if n == 1 {
		return [][]string{[]string{"Q"}}
	}

	str := ""
	for k := 0; k < n; k++ {
		str += "."
	}

	for i := 0; i < len(input); i++ {
		result := []string{}
		for j := 0; j < n; j++ {
			index := input[i][j]
			if index < n-1 && index > 0 {
				result = append(result, str[:index]+"Q"+str[index+1:])
			}
			if index+1 == n {
				result = append(result, str[:index]+"Q")
			}
			if index == 0 {
				result = append(result, "Q"+str[index+1:])
			}
		}
		results = append(results, result)
	}

	return
}
