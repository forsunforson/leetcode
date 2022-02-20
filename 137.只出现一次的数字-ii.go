/*
 * @lc app=leetcode.cn id=137 lang=golang
 *
 * [137] 只出现一次的数字 II
 */
package main

// @lc code=start
func singleNumber(nums []int) int {
	bits := make([]int32, 32)
	for _, num := range nums {
		num := int32(num)
		for i := 0; i < 32; i++ {
			if num == 0 {
				break
			}
			if num&1 == 1 {
				bits[i]++
			}
			num >>= 1
		}
	}
	for i := range bits {
		bits[i] %= 3
	}
	num := int32(0)
	for i := range bits {
		num |= bits[i] << i
	}
	return int(num)
}

// @lc code=end
