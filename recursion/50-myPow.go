package week03

/*
50.Pow(x,y)
实现 pow(x, n) ，即计算 x 的整数 n 次幂函数（即，xn ）。

示例 1：

输入：x = 2.00000, n = 10
输出：1024.00000
示例 2：

输入：x = 2.10000, n = 3
输出：9.26100
示例 3：

输入：x = 2.00000, n = -2
输出：0.25000
解释：2-2 = 1/22 = 1/4 = 0.25


提示：

-100.0 < x < 100.0
-2^31 <= n <= 2^31-1
n 是一个整数
要么 x 不为零，要么 n > 0 。
-10^4 <= x^n <= 10^4
*/

// 位运算+快速幂+迭代
func myPow(x float64, n int) float64 {
	res := 1.0
	if n == 0 {
		return res
	}

	if n < 0 {
		x = 1 / x
		n = -n
	}

	for n > 0 {
		if n&1 == 1 {
			res *= x
		}
		n >>= 1
		x *= x
	}

	return res
}

// 递归超出内存限制
// func myPow(x float64, n int) float64 {
//     res := 1.0
//     if n == 0 {
//         return res
//     }

//     if n < 0 {
//         n = -n
//         x = 1/x
//     }

//     return myPow(x,n-1)*x
// }
