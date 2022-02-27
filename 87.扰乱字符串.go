/*
 * @lc app=leetcode.cn id=87 lang=golang
 *
 * [87] 扰乱字符串
 */
package main

// @lc code=start
func isScramble(s1 string, s2 string) bool {
	// 分解子问题isScramble(s1[:n], s2[:n]) || isScramble(s1[n:], s2[n])
	// dp + dfs dp[len][len][len+1] 分表表示i1, i2和长度
	l := len(s1)
	dp := make([][][]int8, l)
	for i := range dp {
		dp[i] = make([][]int8, l)
		for j := range dp[i] {
			dp[i][j] = make([]int8, l+1)
			for k := range dp[i][j] {
				// 本来用bool也可以，但是为了区分是否初始化
				dp[i][j][k] = -1
			}
		}
	}
	var dfs func(int, int, int) int8
	dfs = func(i1, i2, l int) (res int8) {
		if dp[i1][i2][l] != -1 {
			return dp[i1][i2][l]
		}
		d := &dp[i1][i2][l]
		defer func() { *d = res }()
		sub1, sub2 := s1[i1:i1+l], s2[i2:i2+l]
		if sub1 == sub2 {
			return 1
		}
		if !countDiff(sub1, sub2) {
			return 0
		}
		for i := 1; i < l; i++ {
			if dfs(i1, i2, i) == 1 && dfs(i1+i, i2+i, l-i) == 1 {
				return 1
			}
			if dfs(i1, i2+l-i, i) == 1 && dfs(i1+i, i2, l-i) == 1 {
				return 1
			}
		}
		return 0
	}
	return dfs(0, 0, l) == 1
}

func countDiff(s1, s2 string) bool {
	count := make([]int, 26)
	for i := range s1 {
		count[s1[i]-'a']++
		count[s2[i]-'a']--
	}
	for i := range count {
		if count[i] != 0 {
			return false
		}
	}
	return true
}

// @lc code=end
