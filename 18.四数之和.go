/*
 * @lc app=leetcode.cn id=18 lang=golang
 *
 * [18] 四数之和
 */
package main

import (
	"fmt"
	"sort"
)

// @lc code=start
func fourSum(nums []int, target int) [][]int {
	ans := make([][]int, 0)
	exist := make(map[string]bool)
	sort.Ints(nums)
	for i := 0; i < len(nums)-3; i++ {
		for j := i + 1; j < len(nums)-2; j++ {
			curTarget := target - nums[i] - nums[j]
			m, n := j+1, len(nums)-1
			for m < n {
				if curTarget == nums[m]+nums[n] {
					newSolution := []int{nums[i], nums[j], nums[m], nums[n]}
					if _, ok := exist[fmt.Sprintf("%v", newSolution)]; !ok {
						ans = append(ans, newSolution)
						exist[fmt.Sprintf("%v", newSolution)] = true
					}

					m++
					n--
				} else if curTarget < nums[m]+nums[n] {
					n--
				} else if curTarget > nums[m]+nums[n] {
					m++
				}
			}
		}
	}
	return ans
}

// @lc code=end
