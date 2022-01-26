/*
 * @lc app=leetcode.cn id=338 lang=golang
 *
 * [338] 比特位计数
 */
package main

// @lc code=start
func countBits(n int) []int {
	return count2(n)
	ret := []int{}
	for i := 0; i <= n; i++ {
		count := 0
		num := i
		for num != 0 {
			if num&1 == 1 {
				count++
			}
			num >>= 1
		}
		ret = append(ret, count)
	}
	return ret
}

func count2(n int) []int {
	ret := make([]int, n+1)
	highBit := 0
	for i := 1; i <= n; i++ {
		if i&(i-1) == 0 {
			highBit = i
		}
		ret[i] = ret[i-highBit] + 1
	}
	return ret
}

// @lc code=end
