/*
 * @lc app=leetcode.cn id=312 lang=golang
 *
 * [312] 戳气球
 */
package main

// @lc code=start
func maxCoins(nums []int) int {
	// dp
	copyNums := []int{1}
	copyNums = append(copyNums, nums...)
	copyNums = append(copyNums, 1)
	n := len(copyNums)
	// i, j用来记录(i, j)区间获得的最大分数
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	// 对于k获得的最大值，即nums[i]*nums[k]*nums[j] + dp[i][k] + dp[k][j]
	// 表示先把(i, k) (k, j)的气球先戳爆，最后戳k
	// 注意遍历顺序，从下到上，从左到右
	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			for k := i + 1; k < j; k++ {
				dp[i][j] = max(dp[i][j], dp[i][k]+dp[k][j]+copyNums[k]*copyNums[i]*copyNums[j])
			}
		}
	}
	return dp[0][n-1]
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// @lc code=end
