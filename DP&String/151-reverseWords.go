package week09

import "strings"

func reverseWords(s string) string {
	if s == "" {
		return s
	}

	str := strings.Trim(s, " ")
	words := strings.Split(str, " ")
	index := 0
	for i := 0; i < len(words); i++ {
		if words[i] == "" {
			continue
		}
		if index != i {
			words[index] = words[i]
		}
		index++
	}

	l, r := 0, index-1
	for l < r {
		words[l], words[r] = words[r], words[l]
		l++
		r--
	}

	return strings.Join(words[:index], " ")
}
