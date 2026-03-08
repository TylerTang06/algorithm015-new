package week09

func isIsomorphic(s string, t string) bool {
	myMap := make(map[byte]byte)
	used := make(map[byte]bool)

	for i := 0; i < len(s); i++ {
		if ch, ok := myMap[s[i]]; ok {
			if ch != t[i] {
				return false
			}
		} else {
			if _, used := used[t[i]]; used {
				return false
			}
			myMap[s[i]] = t[i]
			used[t[i]] = true
		}
	}

	return true
}
