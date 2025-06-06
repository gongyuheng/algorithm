package algorithm

/*
给你一个字符串 s，找到 s 中最长的 回文 子串。

示例 1：

输入：s = "babad"
输出："bab"
解释："aba" 同样是符合题意的答案。
示例 2：

输入：s = "cbbd"
输出："bb"


提示：

1 <= s.length <= 1000
s 仅由数字和英文字母组成
*/

/*
回文，a, bb, aba
*/

/*
简单方法：
遍历每个字符，以字符为回文中心扩展
注意：回文中心可能是单个字符，也可能是两个字符
*/
func longestPalindrome(s string) string {
	var start, end int
	for i := 0; i < len(s); i++ {
		l, r := maxLenString(s, i, i)
		if r-l > end-start {
			start, end = l, r
		}

		l, r = maxLenString(s, i, i+1)
		if r-l > end-start {
			start, end = l, r
		}
	}
	return s[start : end+1]
}

func maxLenString(s string, left, right int) (int, int) {
	for ; left >= 0 && right < len(s) && s[left] == s[right]; left, right = left-1, right+1 {
	}
	return left + 1, right - 1
}
