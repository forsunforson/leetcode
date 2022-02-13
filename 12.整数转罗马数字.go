/*
 * @lc app=leetcode.cn id=12 lang=golang
 *
 * [12] 整数转罗马数字
 */
package main

// @lc code=start

type pair struct {
	Val int
	Str string
}

func intToRoman(num int) string {
	// 关键在于尽量选大的数字
	// 数字从大到小
	// 对于map遍历来说，每次并不是按初始化顺序进行的
	mapping := []pair{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}
	ans := ""
	for num > 0 {
		target := ""
		for _, p := range mapping {
			if p.Val <= num {
				target = p.Str
				num -= p.Val
				break
			}
		}
		ans += target
	}
	return ans
}

// @lc code=end
