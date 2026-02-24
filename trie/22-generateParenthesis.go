package week07

// Recursion
func generateParenthesis(n int) []string {
	if n <= 0 {
		return []string{}
	}

	return generateParenthesisRec(n, 0, 0, "", []string{})
}

func generateParenthesisRec(n, left, right int, cur string, res []string) []string {
	if left == right && left == n {
		res = append(res, cur)
		return res
	}

	if left < n {
		res = generateParenthesisRec(n, left+1, right, cur+"(", res)
	}
	if right < left {
		res = generateParenthesisRec(n, left, right+1, cur+")", res)
	}

	return res
}
