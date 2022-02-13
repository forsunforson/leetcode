/*
 * @lc app=leetcode.cn id=354 lang=golang
 *
 * [354] 俄罗斯套娃信封问题
 */
package main

import "sort"

// @lc code=start

func maxEnvelopes(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		return envelopes[i][0] < envelopes[j][0] ||
			envelopes[i][0] == envelopes[j][0] && envelopes[i][1] > envelopes[j][1]
	})
	f := make([]int, 0)
	for _, e := range envelopes {
		h := e[1]
		// 对于整理完后的信封，w是递增的
		// 需要获得一个h也是严格单调递增的子序列
		// 对于w相同的信封，h是递减的，所以替换也不会影响序列长度
		if i := sort.SearchInts(f, h); i < len(f) {
			f[i] = h
		} else {
			f = append(f, h)
		}
	}
	return len(f)
}

// 动态规划解法
func solution1(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		return envelopes[i][0] < envelopes[j][0] ||
			envelopes[i][0] == envelopes[j][0] && envelopes[i][1] > envelopes[j][1]
	})
	n := len(envelopes)
	dp := make([]int, n)
	for i := range dp {
		dp[i] = 1
	}
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			// 如果j能套进i里，dp[i] = dp[j]+1
			if envelopes[i][1] > envelopes[j][1] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	count := 1
	for _, c := range dp {
		count = max(count, c)
	}
	return count
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end
