# 学习笔记

## 动态规划

- 重复性
    - 最优子结构 opt[n] = best of { opt[n-1], opt[n-2],...}
- 定义状态数组
    - 存储中间状态 opt[i], 将作为下一步的依据
- DP方程
    - 递推公式 比如斐布拉切问题：opt[n] = opt[n-1] + opt[n-2]
- 通常中间状态根据具体问题，可能为多维

## 解题技巧

- 64.最短路径和
    - (i,j)可以由(i-1,j)和(i,j-1)过来，因此只需取到这两点路径和小的点过来
    - dp[i][j] = min{dp[i-1][j], dp[i][j-1]} + grid[i][j]
    - https://leetcode-cn.com/problems/minimum-path-sum/
- 91.解码方程
    - 定义dp[i]为s[0...i]的译码方程总和
    - 这里需要分情况讨论，1）如果s[i]='0',dp[i] = dp[i-2]，比如“220” 
    - 2）否则，dp[i] = dp[i-1] + dp[i-2]，比如“121” 
    - 3)这里可以简化，用pre,cur两个变量来存储中间状态，因此空间复杂度可以优化为O(1)
    - https://leetcode-cn.com/problems/decode-ways/
- 221.最大正方形
    - dp[i][j]理解为(i,j)作为正方形的右下点时，此时正方形的边长
    - 此时选出分别以点(i-1,j),(i,j-1),(i-1,j-1)作为右下点的正方形最小边，再加1便是dp[i][j]
    - dp[i][j] = min{d[i-1][j], dp[i][j-1], dp[i-1][j-1]} + 1
    - https://leetcode-cn.com/problems/maximal-square/
- 621.任务调度器
    - 首先统计不同种类的任务，并按需要执行的次数排序，首先安排需要执行次数最多的任务，此时至少需要(max-1)*(n+1)+1的时间
    - 而此时有(max-1)*n的空闲时间，可以安排给其他的任务
    - 此时，如果有任务A，B需要执行次数相等，并且均为max，此时我们需要额外增加一个单位时间
    - 最后，空闲时间片分完，仍还有未执行的任务，则此种情况，需要len(tasks)的时间
    - https://leetcode-cn.com/problems/task-scheduler/
- 647.回文子串
    - dp[i,j]为s[i:j]的子串是否为回文字符串
    - dp[i,i] = true
    - dp[i,j] = dp[i+1,j-1] if s[i] == s[j]
    - j-i < 3, dp[i,j] = true
    - https://leetcode-cn.com/problems/palindromic-substrings/
- ....(待续)

## 总结

- 打破自己思维习惯，学会找重复性
- 定位状态方程(这一步往往最难)
- 不要人肉递归
- 习惯自底向上循环
- 学会优化中间状态存储空间
- 多练习，反复练习