/*
 * @lc app=leetcode.cn id=581 lang=golang
 *
 * [581] 最短无序连续子数组
 */
package main

import (
	"fmt"
	"math"
	"sort"
)

// @lc code=start
func findUnsortedSubarray(nums []int) int {
	// 将数组分为三个子数组A B C
	// 其中A的最大值小于B中任意数；C的最小值大于B中任意数
	start, end := -1, -1
	maxA, minC := math.MinInt32, math.MaxInt32
	for i := 0; i < len(nums); i++ {
		// 从左边最大选右边界
		if nums[i] >= maxA {
			// 如果是左边的最大值，则不更新边界
			maxA = nums[i]
		} else {
			// 比左边最大值小，才更新边界
			// 因为循环从左往右，比左边的最大值小，肯定需要排序
			end = i
		}

		if nums[len(nums)-i-1] <= minC {
			minC = nums[len(nums)-i-1]
		} else {
			start = len(nums) - i - 1
		}
	}
	if end == -1 {
		return 0
	}
	return end - start + 1
}

func solution1(nums []int) int {
	// 暴力解法 nlogn 先排序再比较
	sortedNums := append([]int{}, nums...)
	sort.Ints(sortedNums)
	length := len(nums)
	start, end := length, -1
	for n := 0; n < length; n++ {
		if nums[n] != sortedNums[n] {
			if start > n {
				start = n
			}
			if end < n {
				end = n
			}
		}
	}
	if start == length || end == -1 {
		return 0
	}
	fmt.Printf("start %d end %d", start, end)
	return end - start + 1
}

// @lc code=end
