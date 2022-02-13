/*
 * @lc app=leetcode.cn id=260 lang=golang
 *
 * [260] 只出现一次的数字 III
 */

// @lc code=start
func singleNumber(nums []int) []int {
	// 进行异或计算后，出现两次的元素会抵消
	xorResult := 0
	for _, num := range nums {
		xorResult ^= num
	}
	lowBit := xorResult & -xorResult
	n1, n2 := 0, 0
	for _, num := range nums {
		if num&lowBit == 0 {
			n1 ^= num
		} else {
			n2 ^= num
		}
	}
	return []int{n1, n2}
}

// @lc code=end

