/*
 * @lc app=leetcode.cn id=166 lang=golang
 *
 * [166] 分数到小数
 */
package main

import "strconv"

// @lc code=start
func fractionToDecimal(numerator int, denominator int) string {
	ans := ""
	if numerator == 0 {
		return "0"
	}
	if (numerator >= 0) != (denominator > 0) {
		ans += "-"
	}
	numerator = abs(numerator)
	denominator = abs(denominator)
	ans += strconv.Itoa(numerator / denominator)
	numerator %= denominator
	if numerator == 0 {
		return ans
	}
	ans += "."
	ansMap := map[int]int{}
	for numerator != 0 && ansMap[numerator] == 0 {
		ansMap[numerator] = len(ans)
		numerator *= 10
		ans += strconv.Itoa(numerator / denominator)
		numerator %= denominator
	}
	if numerator != 0 {
		idx := ansMap[numerator]
		tmp := ans[:idx] + "(" + ans[idx:] + ")"
		ans = tmp
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// @lc code=end
