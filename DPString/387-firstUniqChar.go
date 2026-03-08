package week09

func firstUniqChar(s string) int {
	if s == "" {
		return -1
	}

	arr := make([]int, 26)
	for _, c := range s {
		arr[c-'a']++
	}

	for i := 0; i < len(s); i++ {
		if arr[s[i]-'a'] == 1 {
			return i
		}
	}

	return -1
}
