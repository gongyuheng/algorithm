package algorithm

/*
给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
例如，121 是回文，而 123 不是。


示例 1：

输入：x = 121
输出：true
示例 2：

输入：x = -121
输出：false
解释：从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
示例 3：

输入：x = 10
输出：false
解释：从右向左读, 为 01 。因此它不是一个回文数。


提示：

-231 <= x <= 231 - 1

进阶：你能不将整数转为字符串来解决这个问题吗？
*/

// 自己算的，完全翻转+比较
// func isPalindrome(x int) bool {
// 	if x < 0 {
// 		return false
// 	}

// 	before, after := x, 0
// 	for before > 0 {
// 		after = after*10 + before%10
// 		before /= 10
// 	}
// 	return after == x
// }

// 官方题解，翻转一半就可以了, 确实牛逼
func isPalindrome(x int) bool {
	// 特殊处置
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	revertedNumber := 0
	// 一直翻转直到翻转后的数 >= x
	for x > revertedNumber {
		revertedNumber = revertedNumber*10 + x%10
		x /= 10
	}

	// 回文的两种情况 1221 12321 即 12 == 12 || 12 == 123/10
	return x == revertedNumber || x == revertedNumber/10
}
