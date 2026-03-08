package week08

func totalNQueens(n int) int {
	if n < 1 {
		return 0
	}

	return totalNQueensUseBits(n, 0, 0, 0, 0, 0)
}

func totalNQueensUseBits(n, row, cols, sum, diff, count int) int {
	if row >= n {
		count++
		return count
	}

	bits := (^(cols | sum | diff)) & ((1 << n) - 1) // 得到当前所有空位
	for bits != 0 {
		p := bits & -bits // 取到最低位 1
		count = totalNQueensUseBits(n, row+1, cols|p, (sum|p)<<1, (diff|p)>>1, count)
		bits = bits & (bits - 1) // 去掉最低位 1
	}

	return count
}
