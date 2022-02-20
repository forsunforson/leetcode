/*
 * @lc app=leetcode.cn id=386 lang=golang
 *
 * [386] 字典序排数
 */
package main

// @lc code=start
func lexicalOrder(n int) []int {
	ans := make([]int, n)
	// 迭代，可以不用O(n)额外空间
	num := 1
	idx := 0
	for idx < n {
		for num <= n {
			ans[idx] = num
			idx++
			num *= 10
		}
		for num%10 == 9 || num > n {
			num /= 10
		}
		num++
	}
	return ans
}

// @lc code=end
// 递归
func recurse(n int) []int {
	ans := []int{}
	var dfs func(int)
	dfs = func(cur int) {
		if cur > n {
			return
		}
		ans = append(ans, cur)
		for i := 0; i <= 9; i++ {
			dfs(cur*10 + i)
		}
	}
	for i := 1; i <= n && i <= 9; i++ {
		dfs(i)
	}
	return ans
}
