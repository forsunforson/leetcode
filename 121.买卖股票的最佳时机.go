/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 */
package main

// @lc code=start
func maxProfit(prices []int) int {
	// 拿俩数组记录左边最小值和右边最大值
	left := make([]int, len(prices))
	right := make([]int, len(prices))
	min := prices[0]
	max := prices[len(prices)-1]
	for i := 0; i < len(prices); i++ {
		if prices[i] <= min {
			min = prices[i]
		}
		left[i] = min
		if prices[len(prices)-1-i] >= max {
			max = prices[len(prices)-1-i]
		}
		right[len(prices)-1-i] = max
	}
	ret := 0
	for i := 0; i < len(prices); i++ {
		cur := right[i] - left[i]
		if cur > ret {
			ret = cur
		}
	}
	return ret
}

func maxProfit2(prices []int) int {
	min := prices[0]
	ret := 0
	for _, price := range prices {
		if price < min {
			min = price
		} else {
			profit := price - min
			if profit > ret {
				ret = profit
			}
		}
	}
	return ret
}

// @lc code=end
