/*
 * @lc app=leetcode.cn id=13 lang=golang
 *
 * [13] 罗马数字转整数
 */
package main

// @lc code=start
func romanToInt(s string) int {
	mapping := map[string]int{
		"M":  1000,
		"CM": 900,
		"D":  500,
		"CD": 400,
		"C":  100,
		"XC": 90,
		"L":  50,
		"XL": 40,
		"X":  10,
		"IX": 9,
		"V":  5,
		"IV": 4,
		"I":  1,
	}
	ans := 0
	idx := 0
	for idx < len(s) {
		if idx+2 <= len(s) {
			target := s[idx : idx+2]
			if v, ok := mapping[target]; ok {
				ans += v
				idx += 2
				continue
			}
		}

		target := s[idx : idx+1]
		v := mapping[target]
		ans += v
		idx += 1

	}
	return ans
}

// @lc code=end
