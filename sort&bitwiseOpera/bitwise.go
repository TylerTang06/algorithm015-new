package week08

// XOR
// x ^ 0 = x
// x ^ 1s = ~x // 1s = ~0
// x ^ (~x) = 1s
// x ^ x = 0
// a ^ b = c, a ^ c = b, b ^ c = a // swap
// a ^ b ^ c = a ^ (b ^ c) = (a ^ b) ^ c // associative
func XOR(a, b int) int {
	return a ^ b
}

// Clean n bits to 0 on the right
func ToZeroRight(x, n int) int {
	return x & (^int(0) << n)
}

// get value of the nth bit
func GetBitOfN(x, n int) int {
	return (x >> n) & 1
}

// get power value of the nth bit
func GetPowerOfN(x, n int) int {
	return x & (1 << (n - 1))
}

// update value of the nth bit to 1
func UpdateToOneOfN(x, n int) int {
	return x | (1 << n)
}

// update value of the nth bit to 0
func UpdateToZeroOfN(x, n int) int {
	return x & (^(1 << n))
}

// update value from nth to mth bits to 0,
// mth is the Hightest bit
func UpdateToZeroFromN(x, n int) int {
	return x & ((1 << n) - 1)
}

// upate value from 0th to nth bits
func UpdateToZeroToN(x, n int) int {
	return x & (^((1 << (n + 1)) - 1))
}

// clean the first 1 on the right
func ToZeroFirstOneRight(x int) int {
	return x & (x - 1)
}

// get the first 1 on the right
func GetFirstOneRight(x int) int {
	return x & -x
}

//  Number of 1 Bits
func HammingWeight(num uint32) int {
	times := 0
	for num != 0 {
		num = num & (num - 1)
		times++
	}

	return times
}

// Power of Two
func IsPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}

	return (n & (n - 1)) == 0
}

// Given a non negative integer number num.
// For every numbers i in the range 0 ≤ i ≤ num
// calculate the number of 1's in their binary representation
// and return them as an array.
func CountBits(num int) []int {
	counts := make([]int, num+1)
	for i := 1; i <= num; i++ {
		counts[i] += counts[i&(i-1)] + 1
	}

	return counts
}
