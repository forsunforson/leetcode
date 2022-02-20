/*
 * @lc app=leetcode.cn id=718 lang=golang
 *
 * [718] 最长重复子数组
 */
package main

// @lc code=start
func findLength(nums1 []int, nums2 []int) int {
	// dp 最长公共子串
	dp := make([][]int, len(nums1)+1)
	for i := range dp {
		dp[i] = make([]int, len(nums2)+1)
	}
	ans := 0
	for i := len(nums1) - 1; i >= 0; i-- {
		for j := len(nums2) - 1; j >= 0; j-- {
			if nums1[i] == nums2[j] {
				dp[i][j] = dp[i+1][j+1] + 1
				ans = max(ans, dp[i][j])
			}
		}
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end
