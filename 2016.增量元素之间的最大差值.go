/*
 * @lc app=leetcode.cn id=2016 lang=golang
 *
 * [2016] 增量元素之间的最大差值
 */
package main

// @lc code=start
func maximumDifference(nums []int) int {
	ret := -1
	curMin := nums[0]
	for _, num := range nums {
		curMin = min(curMin, num)
		if num > curMin {
			ret = max(ret, num-curMin)
		}
	}
	return ret
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// @lc code=end
