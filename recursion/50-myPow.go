package week03

func myPow(x float64, n int) float64 {
	res := 1.0
	if n == 0 {
		return res
	}

	if n < 0 {
		x = 1 / x
		n = -n
	}

	for n > 0 {
		if n&1 == 1 {
			res *= x
		}
		n >>= 1
		x *= x
	}

	return res
}
