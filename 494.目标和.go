/*
 * @lc app=leetcode.cn id=494 lang=golang
 *
 * [494] 目标和
 */
package main

// @lc code=start
func findTargetSumWays(nums []int, target int) int {
	return findTargetSumWays2(nums, target)
}

// 最简单的方法，暴力枚举
func findTargetSumWays1(nums []int, target int) int {
	count := 0
	find1Helper(0, 0, target, nums, &count)
	return count
}

func find1Helper(pre, curIdx, target int, nums []int, count *int) {
	if curIdx >= len(nums) {
		return
	}

	curAdd := pre + nums[curIdx]
	if curIdx == len(nums)-1 && curAdd == target {
		*count += 1
	}
	find1Helper(curAdd, curIdx+1, target, nums, count)
	curSub := pre - nums[curIdx]
	if curIdx == len(nums)-1 && curSub == target {
		*count += 1
	}
	find1Helper(curSub, curIdx+1, target, nums, count)
}

// DP
func findTargetSumWays2(nums []int, target int) int {
	// n个数字中，选择x个减去
	sum := 0
	for _, num := range nums {
		sum += num
	}
	n := len(nums)
	diff := sum - target
	if diff < 0 {
		return 0
	}
	neg := diff / 2
	if diff%2 == 1 {
		return 0
	}
	// 2维数组用来记录前i个数字获得j结果的count数
	dp := make([][]int, len(nums)+1)
	for i := range dp {
		dp[i] = make([]int, neg+1)
	}
	dp[0][0] = 1
	// 初始边界，i=0
	for i, num := range nums {
		for j := 0; j <= neg; j++ {
			dp[i+1][j] = dp[i][j]
			if j >= num {
				dp[i+1][j] += dp[i][j-num]
			}
		}
	}

	return dp[n][neg]
}

// @lc code=end
