package algorithm

/*
给定一个字符串 s ，请你找出其中不含有重复字符的 最长 子串 的长度。

示例 1:

输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:

输入: s = "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: s = "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
*/

func LengthOfLongestSubstring(s string) int {
	m := make(map[byte]int) // 根据情况可以考虑使用数组映射

	// 使用滑动窗口 [l,r)
	var maxlen int
	for l, r := 0, 0; l < len(s); l++ {
		// 右游标一直到重复的位置
		for r < len(s) && m[s[r]] == 0 {
			m[s[r]]++
			r++
		}
		maxlen = max(maxlen, r-l)

		// 左游标移动一位
		m[s[l]]--
	}
	return maxlen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
