package week09

func reverseWords1(s string) string {
	if s == "" || s == " " {
		return s
	}

	b := []byte(s)
	le := 0
	for i := 0; i < len(b); i++ {
		if b[i] != ' ' {
			le++
		}
		if b[i] == ' ' || len(b)-1 == i {
			l, r := i-le, i-1
			if len(b)-1 == i {
				l, r = i-le+1, i
			}
			for l < r {
				b[l], b[r] = b[r], b[l]
				l++
				r--
			}
			le = 0
		}
	}

	return string(b)
}
