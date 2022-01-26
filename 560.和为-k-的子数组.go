/*
 * @lc app=leetcode.cn id=560 lang=golang
 *
 * [560] 和为 K 的子数组
 */
package main

// @lc code=start
func subarraySum(nums []int, k int) int {
	return subarraySum2(nums, k)
}

func subarraySum1(nums []int, k int) int {
	// n^2 的复杂度
	count := 0
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum == k {
				count++
			}
		}
	}
	return count
}

// 前缀和数组用来记录，避免重复计算
func subarraySum2(nums []int, k int) int {
	count := 0
	pre := 0                    // 用来记录前缀和
	sumMap := make(map[int]int) // 用来记录出现的次数
	sumMap[0] = 1
	for i := 0; i < len(nums); i++ {
		pre += nums[i]
		if c, ok := sumMap[pre-k]; ok {
			// 为什么需要记录pre - k的次数？
			// 对于pre来说是sum[0,i]，减去sum[0,j-1]，就获得了对应sum[j-1,i]出现的次数
			count += c
		}
		sumMap[pre] += 1
	}
	return count
}

// @lc code=end
