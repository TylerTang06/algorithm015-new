package week07

func solveNQueens(n int) [][]string {
	if n < 1 {
		return [][]string{}
	}

	results := solveNQueensRec(n, 0, map[int]bool{}, map[int]bool{}, map[int]bool{}, []int{}, [][]int{})

	return generateNQueensResults(n, results)
}

func solveNQueensRec(n, row int, colMap, sumMap, diffMap map[int]bool, curResult []int, results [][]int) [][]int {
	if row == n {
		results = append(results, append([]int{}, curResult[:]...))
		return results
	}

	for col := 0; col < n; col++ {
		if _, ok := colMap[col]; ok {
			continue
		}
		if _, ok := sumMap[row+col]; ok {
			continue
		}
		if _, ok := diffMap[row-col]; ok {
			continue
		}

		colMap[col], sumMap[row+col], diffMap[row-col] = true, true, true
		newCurResult := append(append([]int{}, curResult[:]...), col)
		results = solveNQueensRec(n, row+1, colMap, sumMap, diffMap, newCurResult, results)
		delete(colMap, col)
		delete(sumMap, row+col)
		delete(diffMap, row-col)
	}

	return results
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
