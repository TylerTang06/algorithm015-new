package week02

/*
242.有效的字母异位词
给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的 字母异位词。



示例 1:

输入: s = "anagram", t = "nagaram"
输出: true
示例 2:

输入: s = "rat", t = "car"
输出: false


提示:

1 <= s.length, t.length <= 5 * 10^4
s 和 t 仅包含小写字母
*/
func isAnagram(s string, t string) bool {
	myMap := map[rune]int{}
	for _, str := range s {
		if _, ok := myMap[str]; ok {
			myMap[str]++
		} else {
			myMap[str] = 1
		}
	}

	for _, str := range t {
		if _, ok := myMap[str]; ok {
			myMap[str]--
			if myMap[str] == 0 {
				delete(myMap, str)
			}
		} else {
			return false
		}
	}
	if len(myMap) != 0 {
		return false
	}

	return true
}
