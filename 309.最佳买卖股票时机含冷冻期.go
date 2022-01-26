/*
 * @lc app=leetcode.cn id=309 lang=golang
 *
 * [309] 最佳买卖股票时机含冷冻期
 */
package main

// @lc code=start
func maxProfit(prices []int) int {
	l := len(prices)
	// 根据是否在冷冻期，是否买卖，把当前状态分为0 1 2
	// 0 持有股票 1 不持有股票、不冷冻 2 不持有股票、冷冻
	f0, f1, f2 := -prices[0], 0, 0
	for i := 1; i < l; i++ {
		// 当前持有股票，那就是从上个不冷冻期买的，或者一直没卖过
		f0 = max(f1-prices[i], f0)
		// 当前不持有股票且不冷冻，只能是已经卖掉股票了，冷冻不冷冻都一样
		f1 = max(f1, f2)
		// 当前不持有股票且冷冻，一定是从持有股票中卖掉了
		f2 = f0 + prices[i]
	}
	return max(f1, f2)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end
