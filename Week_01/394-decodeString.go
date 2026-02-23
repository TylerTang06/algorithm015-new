package week01

import "container/list"

/*
394.字符串解码

给定一个经过编码的字符串，返回它解码后的字符串。

编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。

你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。

此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像 3a 或 2[4] 的输入。

测试用例保证输出的长度不会超过 10^5。

示例 1：

输入：s = "3[a]2[bc]"
输出："aaabcbc"
示例 2：

输入：s = "3[a2[c]]"
输出："accaccacc"
示例 3：

输入：s = "2[abc]3[cd]ef"
输出："abcabccdcdcdef"
示例 4：

输入：s = "abc3[cd]xyz"
输出："abccdcdcdxyz"


提示：

1 <= s.length <= 30
s 由小写英文字母、数字和方括号 '[]' 组成
s 保证是一个 有效 的输入。
*/

func decodeString(s string) string {
	var res string
	var num int
	stk := list.New()
	for _, v := range s {
		switch v {
		case '[':
			stk.PushBack(num)
			stk.PushBack(res)
			res, num = "", 0
		case ']':
			resOld := stk.Back().Value.(string)
			stk.Remove(stk.Back())
			numOld := stk.Back().Value.(int)
			stk.Remove(stk.Back())
			res = resOld + duplicateString(numOld, res)
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			num = num*10 + int(v-'0')
		default:
			res += string(v)
		}

	}
	return res
}

func duplicateString(n int, str string) string {
	var res string
	for i := 0; i < n; i++ {
		res += str
	}

	return res
}
