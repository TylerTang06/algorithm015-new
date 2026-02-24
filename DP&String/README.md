# 学习笔记

## 动态规划

- 重复性
    - 最优子结构 opt[n] = best of { opt[n-1], opt[n-2],...}
- 定义状态数组
    - 存储中间状态 opt[i], 将作为下一步的依据
- DP方程
    - 递推公式 比如斐布拉切问题：opt[n] = opt[n-1] + opt[n-2]
- 通常中间状态根据具体问题，可能为多维

## 动态规划-解题技巧

- 72.编辑距离
    - dp[i][j]，表示第一个字符串的前i个字符与第二个字符串的前j个字符的编辑距离
    - dp[i][j] = min{dp[i-1][j], dp[i][j-1], dp[i-1][j-1]} + 1 if word1[i-1] != word2[j-1]
        else dp[i][j] = dp[i-1][j-1]
    - https://leetcode-cn.com/problems/edit-distance/
- 1143.最长公共子序列
    - dp[i][j]，表示第一个字符串的前i个字符与第二个字符串的前j个字符的最长公共子序列数
    - dp[i][j] = max{dp[i-1][j], dp[i][j-1]} if text1[i-1] != text2[j-1]
        else dp[i][j] = dp[i-1][j-1] + 1
    - https://leetcode-cn.com/problems/longest-common-subsequence/
- 1143.最长公共子序列
    - dp[i][j]，表示第i个到第j的子串是否是回文子串
    - dp[i][j] = dp[i+1][j-1] if s[i] == s[j]
        else dp[i][j] = false
    - https://leetcode-cn.com/problems/longest-palindromic-substring/
- 10.正则表达式匹配
    - ...

## 字符串算法-解题技巧

- 使用特定数据结构存储统计信息
    - 使用map，如字符串中的第一个唯一字符，异位词问题
    - 使用Trie，如最长公共前缀
- 将字符串转换为[]byte数组
    - string 在golang中是不可变的
    - 利于交换操作，如反转字符串

    ```golang
    b := []byte(s)
    ```
- 双指针操作，如反转字符串系列题，回文子串系列题
- 中心往两边扩散，如回文子串
- 子串和子序列，使用DP解决定义：dp[i][j]，表第一个字符串的前i个字符与第二个字符串的前j个字符的状态
    - 子串要求必须是连续的，子序列只需保证不改变字符序列
- 使用状态机解决字符串转换整数 (atoi) 问题
    - 状态机的定义，以及状态转移的定义
    ```golang
    // using state machine to do it
    var states = map[byte][]byte{
        // preState: " ", +/-, NUMBER, other
        START:  []byte{START, SIGNED, NUMBER, END},
        SIGNED: []byte{END, END, NUMBER, END},
        NUMBER: []byte{END, END, NUMBER, END},
        END:    []byte{END, END, END, END},
    }

    type AtoiStateMachine struct {
        ans   int
        sign  int // 1 or -1
        state byte
    }
    ```
    - 注意与最大和最小整数进行边界判断
    ```golang
    pop := sm.sign * int(c)
    if sm.sign == 1 && (sm.ans > math.MaxInt32/10 || sm.ans == math.MaxInt32/10 && pop > 7) {
        sm.state = END
        sm.ans = math.MaxInt32
        return sm.ans, false
    }
    if sm.sign == -1 && (sm.ans < math.MinInt32/10 || sm.ans == math.MinInt32/10 && pop < -8) {
        sm.state = END
        sm.ans = math.MinInt32
        return sm.ans, false
    }
    ```
