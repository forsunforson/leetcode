/*
 * @lc app=leetcode.cn id=16 lang=golang
 *
 * [16] 最接近的三数之和
 */
package main

import "sort"

// @lc code=start
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	sum := nums[0] + nums[1] + nums[2]
	minGap := abs(target, sum)
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				if abs(target, nums[i]+nums[j]+nums[k]) < minGap {
					sum = nums[i] + nums[j] + nums[k]
					minGap = abs(target, nums[i]+nums[j]+nums[k])
				}
			}
		}
	}
	return sum
}

func abs(x, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}

// @lc code=end
