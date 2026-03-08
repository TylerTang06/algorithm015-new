package week04

func isPerfectSquare(num int) bool {
	if num <= 0 {
		return false
	}

	x := 1
	// eg. 16 = 1 + 3 + 5 + 7
	for num > 0 {
		num -= x
		x += 2
	}

	return num == 0
}
