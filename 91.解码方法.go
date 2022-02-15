/*
 * @lc app=leetcode.cn id=91 lang=golang
 *
 * [91] 解码方法
 */
package main

// @lc code=start
func numDecodings(s string) int {
	// dp
	dp := make([]int, len(s)+1)
	if s[0] == '0' {
		return 0
	}
	dp[0] = 1
	for i := 1; i <= len(s); i++ {
		// dp[i] = dp[i-1]+dp[i-2]
		if s[i-1] != '0' {
			dp[i] += dp[i-1]
		}
		if i > 1 && s[i-2] != '0' {
			// 使用两个字符
			if (s[i-2]-'0')*10+(s[i-1]-'0') <= 26 {
				dp[i] += dp[i-2]
			}
		}
	}
	return dp[len(s)]
}

// @lc code=end
func backtrack(s string) int {
	// 太慢了
	count := 0
	var dfs func(int)
	dfs = func(idx int) {
		if idx == len(s) {
			count++
			return
		}
		if s[idx] == '0' {
			return
		}
		if s[idx] > '2' {
			dfs(idx + 1)
		} else {
			if s[idx] == '1' {
				if idx+1 < len(s) {
					dfs(idx + 1)
					dfs(idx + 2)
				} else {
					dfs(idx + 1)
				}

			} else if idx+1 < len(s) && s[idx+1] <= '6' {
				dfs(idx + 1)
				dfs(idx + 2)
			} else {
				dfs(idx + 1)
			}
		}
	}
	dfs(0)
	return count
}
