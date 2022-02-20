/*
 * @lc app=leetcode.cn id=189 lang=golang
 *
 * [189] 轮转数组
 */
package main

// @lc code=start
func rotate(nums []int, k int) {
	k = k % len(nums)
	reverse(nums, 0, len(nums)-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, len(nums)-1)
}

func reverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}

}

// @lc code=end
func sol1(nums []int, k int) {
	k = k % len(nums)
	ans := make([]int, len(nums))
	for i := range nums {
		ans[(i+k)%len(nums)] = nums[i]
	}
	for i := range nums {
		nums[i] = ans[i]
	}
}

func sol2(nums []int, k int) {
	k = k % len(nums)
	for i := 0; i < k; i++ {
		tmp := nums[len(nums)-1]
		for j := len(nums) - 1; j > 0; j-- {
			nums[j] = nums[j-1]
		}
		nums[0] = tmp
	}
}
