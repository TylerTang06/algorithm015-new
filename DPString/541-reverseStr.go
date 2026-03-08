package week09

func reverseStr(s string, k int) string {
	b := []byte(s)
	flag := 1
	for i := 0; i < len(b); {
		end := len(b)
		if i+k < len(b) {
			end = i + k
		}
		if flag == 1 {
			for j := end - 1; i < j; j-- {
				b[i], b[j] = b[j], b[i]
				i++
			}
		}
		i, flag = end, -flag
	}

	return string(b)
}
