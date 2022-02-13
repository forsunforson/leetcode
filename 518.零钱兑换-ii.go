/*
 * @lc app=leetcode.cn id=518 lang=golang
 *
 * [518] 零钱兑换 II
 */
package main

import "fmt"

// @lc code=start
func change(amount int, coins []int) int {
	// dp[target]
	dp := make([]int, amount+1)
	dp[0] = 1
	// 这种方式会重复计算例如1+2和2+1
	// for i := 1; i < len(dp); i++ {
	// 	for _, coin := range coins {
	// 		if i-coin >= 0 {
	// 			dp[i] += dp[i-coin]
	// 		}
	// 	}
	// }
	// 确保小的硬币取完了，再取大的
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] += dp[i-coin]
		}
	}
	fmt.Printf("%v\n", dp)
	return dp[amount]
}

// @lc code=end
