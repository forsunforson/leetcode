/*
 * @lc app=leetcode.cn id=29 lang=golang
 *
 * [29] 两数相除
 */
package main

import "math"

// @lc code=start
func divide(dividend int, divisor int) int {
	// 本体考虑完边界情况后，还需要满足时间限制
	// 使用二分法查找答案
	if dividend == math.MinInt32 { // 考虑被除数为最小值的情况
		if divisor == 1 {
			return math.MinInt32
		}
		if divisor == -1 {
			return math.MaxInt32
		}
	}
	if divisor == math.MinInt32 { // 考虑除数为最小值的情况
		if dividend == math.MinInt32 {
			return 1
		}
		return 0
	}
	if dividend == 0 { // 考虑被除数为 0 的情况
		return 0
	}

	neg := false
	if dividend > 0 && divisor < 0 {
		divisor = -divisor
		neg = true
	} else if dividend < 0 && divisor > 0 {
		dividend = -dividend
		neg = true
	} else if dividend < 0 && divisor < 0 {
		dividend = -dividend
		divisor = -divisor
	}
	result := 0
	l, r := 1, math.MaxInt32
	for l <= r {
		mid := l + (r-l)>>1
		if add(divisor, mid, dividend) {
			result = mid
			if mid == math.MaxInt32 {
				break
			}
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	if neg {
		result = -result
	}
	return result
}

func add(divisor, mid, dividend int) bool {
	// result 表示mid*divisor的结果
	for result, add := 0, divisor; mid > 0; mid >>= 1 {
		if mid&1 > 0 {
			// 说明mid是奇数
			if result < dividend-add {
				return false
			}
			// 把结果加上
			result += add
		}
		if mid != 1 {
			// 还可以倍增
			// 如果2*add>dividend，说明当前值太大了
			if add < dividend-add {
				return false
			}
			add += add
		}
	}
	return true
}

// @lc code=end
