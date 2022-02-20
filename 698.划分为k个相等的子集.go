/*
 * @lc app=leetcode.cn id=698 lang=golang
 *
 * [698] 划分为k个相等的子集
 */
package main

import "sort"

// @lc code=start
func canPartitionKSubsets(nums []int, k int) bool {
	// 对于划分子集的问题，一般是穷举或者dp
	// 数组里的数是否选择由bit表示
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%k != 0 {
		return false
	}
	target := sum / k
	sort.Ints(nums)
	if nums[len(nums)-1] > target {
		return false
	}
	lens := len(nums)
	size := 1 << lens
	dp := make([]bool, size)
	dp[0] = true
	curSum := make([]int, size)
	for i := 0; i < size; i++ {
		if !dp[i] {
			continue
		}
		// 将数字一个一个往里加
		for j := 0; j < lens; j++ {
			// i记录着当前用过的数字，j必须是没用过的
			if i&(1<<j) != 0 {
				continue
			}
			next := i | (1 << j)
			if dp[next] {
				// 如果访问过，跳过
				continue
			}
			if (curSum[i]%target)+nums[j] <= target {
				curSum[next] = curSum[i] + nums[j]
				dp[next] = true
			} else {
				break
			}
		}
	}
	return dp[size-1]
}

// @lc code=end
