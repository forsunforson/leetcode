/*
 * @lc app=leetcode.cn id=90 lang=golang
 *
 * [90] 子集 II
 */
package main

import (
	"sort"
)

// @lc code=start
func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	ans := [][]int{}
	var dfs func(bool, int)
	path := []int{}
	// 用bool剪枝
	dfs = func(choosePre bool, idx int) {
		if idx == len(nums) {
			tmp := append([]int{}, path...)
			ans = append(ans, tmp)
			return
		}
		// 不选当前值
		dfs(false, idx+1)
		// 选当前值
		if !choosePre && idx > 0 && nums[idx-1] == nums[idx] {
			// 当前不选，但是下一个一样，说明会重复
			// 已经在“当前选”的分支里出现过了
			return
		}
		path = append(path, nums[idx])
		dfs(true, idx+1)
		path = path[:len(path)-1]
	}
	dfs(false, 0)
	return ans
}

// @lc code=end
