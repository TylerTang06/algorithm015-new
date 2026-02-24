package week09

import "unicode"

func reverseOnlyLetters(S string) string {
	if S == "" || S == " " {
		return S
	}

	b := []byte(S)
	l, r := 0, len(b)-1
	for l < r {
		if !unicode.IsLetter(rune(b[l])) {
			l++
			continue
		}
		if !unicode.IsLetter(rune(b[r])) {
			r--
			continue
		}
		b[l], b[r] = b[r], b[l]
		l++
		r--
	}

	return string(b)
}
