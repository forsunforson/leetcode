/*
 * @lc app=leetcode.cn id=89 lang=golang
 *
 * [89] 格雷编码
 */
package main

// @lc code=start
func grayCode(n int) []int {
	// n位gray编码可由n-1位gray编码推出
	// 分别是n-1位正序第一位加0和倒序加1
	if n == 1 {
		return []int{0, 1}
	}
	// n=2, 2;n=3, 4
	ans := []int{}
	last := grayCode(n - 1)
	for _, v := range last {
		ans = append(ans, v)
	}
	plus := 1 << (n - 1)
	for i := len(last) - 1; i >= 0; i-- {
		ans = append(ans, last[i]+plus)
	}
	return ans
}

// @lc code=end
