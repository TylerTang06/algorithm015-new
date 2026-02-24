package week06

import "math"

// pace l, r two pointers
func minWindow(s string, t string) string {
	if s == "" || t == "" {
		return ""
	}

	tMap, curMap := map[byte]int{}, map[byte]int{}
	for i := 0; i < len(t); i++ {
		tMap[t[i]]++
	}

	lPos, rPos, curLen := -1, -1, math.MaxInt32
	for l, r := 0, 0; r < len(s); r++ {
		if tMap[s[r]] > 0 {
			curMap[s[r]]++
		}
		for isHasT(tMap, curMap) && l <= r {
			if r-l+1 < curLen {
				curLen = r - l + 1
				lPos, rPos = l, r
			}
			if _, ok := tMap[s[l]]; ok {
				curMap[s[l]]--
			}
			l++
		}
	}

	if lPos == -1 {
		return ""
	}

	return s[lPos : rPos+1]
}

// for example, t = "AABBC"
func isHasT(tMap, curMap map[byte]int) bool {
	for k, v := range tMap {
		if v1, ok := curMap[k]; !ok || v1 < v {
			return false
		}
	}

	return true
}
