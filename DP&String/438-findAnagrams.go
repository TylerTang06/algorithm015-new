package week09

func findAnagrams(s string, p string) []int {
	if s == "" || p == "" {
		return []int{}
	}

	pLen, pArr := len(p), make([]int, 26)
	for _, c := range p {
		pArr[c-'a']++
	}

	res, usedArr := []int{}, make([]int, 26)
	l, r := 0, 0
	for r < len(s) {
		rVar := s[r] - 'a'
		usedArr[rVar]++
		r++

		for usedArr[rVar] > pArr[rVar] {
			usedArr[s[l]-'a']--
			l++
		}

		if r-l == pLen {
			res = append(res, l)
		}
	}

	return res
}
