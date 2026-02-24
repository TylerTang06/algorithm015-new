package week08

func reverseBits(num uint32) uint32 {
	res := uint32(0)
	for i := 31; i >= 0; i-- {
		res += (num & 1) << i
		num >>= 1
	}

	return res
}
