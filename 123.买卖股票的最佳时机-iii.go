/*
 * @lc app=leetcode.cn id=123 lang=golang
 *
 * [123] 买卖股票的最佳时机 III
 */

// @lc code=start
func maxProfit(prices []int) int {
	buy1, sell1 := -prices[0], 0
	buy2, sell2 := -prices[0], 0
	// 因为一次遍历，所以sell1最小值一定发生再buy2之前
	for i := 1; i < len(prices); i++ {
		// 买的越便宜越好
		buy1 = max(buy1, -prices[i])
		sell1 = max(sell1, prices[i]+buy1)
		buy2 = max(buy2, sell1-prices[i])
		sell2 = max(sell2, buy2+prices[i])
	}
	return sell2
}

// 多重循环不能通过所有用例
func solution1(prices []int) int {
	// dp[time][day] 交易次数和日期构成的矩阵
	n := len(prices)
	dp := make([][]int, 2)
	dp[0] = make([]int, n)
	dp[1] = make([]int, n)
	minPrice := prices[0]
	// 计算交易一次的最大值
	for i := 0; i < n; i++ {
		minPrice = min(minPrice, prices[i])
		dp[0][i] = max(0, prices[i]-minPrice)
	}
	// 第二次交易的最大值基于第一次交易
	for i := 1; i < n; i++ {
		dp[1][i] = max(dp[0][i], dp[1][i-1])
		for j := 0; j < i; j++ {
			minPrice = findMinInRange(prices, j, i)
			dp[1][i] = max(dp[1][i], dp[0][j]+prices[i]-minPrice)
		}
	}
	fmt.Printf("%v\n", dp)
	return dp[1][n-1]
}

func findMinInRange(prices []int, x, y int) int {
	minPrice := prices[x]
	for i := x; i <= y; i++ {
		minPrice = min(minPrice, prices[i])
	}
	return minPrice
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// @lc code=end

