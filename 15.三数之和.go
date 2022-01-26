/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */
package main

import (
	"fmt"
	"sort"
)

// @lc code=start
func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}
	ret := make([][]int, 0)
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j, k := i+1, len(nums)-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == 0 {
				ret = append(ret, []int{nums[i], nums[j], nums[k]})
				fmt.Printf("%d %d %d \n", nums[i], nums[j], nums[k])
				j++
				for j < k && nums[j-1] == nums[j] {
					j++
				}
				k--
				for k > j && nums[k] == nums[k+1] {
					k--
				}
			} else if sum > 0 {
				k--
			} else {
				j++
			}
		}
	}
	return ret
}

// @lc code=end
