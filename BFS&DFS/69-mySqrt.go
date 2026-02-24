package week04

func mySqrt(x int) int {
	if x <= 0 {
		return 0
	}

	xn := x
	for xn*xn > x {
		xn = (xn + x/xn) / 2
	}

	return xn
}
