package week01

import "container/list"

/*
20. 有效的括号

给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
每个右括号都有一个对应的相同类型的左括号。

示例 1：

输入：s = "()"

输出：true

示例 2：

输入：s = "()[]{}"

输出：true

示例 3：

输入：s = "(]"

输出：false

示例 4：

输入：s = "([])"

输出：true

示例 5：

输入：s = "([)]"

输出：false



提示：

1 <= s.length <= 10^4
s 仅由括号 '()[]{}' 组成

*/

var bracketPair = map[rune]rune{
	')': '(',
	'}': '{',
	']': '[',
}

func isValid(s string) bool {
	stk := list.New()

	for _, ss := range s {
		if v, ok := bracketPair[ss]; ok {
			if stk.Len() > 0 && v == stk.Back().Value.(rune) {
				stk.Remove(stk.Back())
			} else {
				return false
			}
		} else {
			stk.PushBack(ss)
		}
	}

	return stk.Len() == 0
}
