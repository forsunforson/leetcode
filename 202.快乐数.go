/*
 * @lc app=leetcode.cn id=202 lang=golang
 *
 * [202] 快乐数
 */
package main

// @lc code=start
func isHappy(n int) bool {
	visited := map[int]bool{}
	for !visited[n] {
		if n == 1 {
			return true
		}
		visited[n] = true
		n = getNext(n)
	}
	return false
}

func getNext(num int) int {
	sum := 0
	for num > 0 {
		n := num % 10
		sum += n * n
		num = num / 10
	}
	return sum
}

// @lc code=end
