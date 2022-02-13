/*
 * @lc app=leetcode.cn id=40 lang=golang
 *
 * [40] 组合总和 II
 */
package main

import (
	"sort"
)

// @lc code=start
func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	ans := make([][]int, 0)
	numCount := make(map[int]int)
	for _, candidate := range candidates {
		numCount[candidate]++
	}
	var backtrack func(int, int, *[]int)
	backtrack = func(target, idx int, path *[]int) {
		if target == 0 {
			tmp := append([]int{}, *path...)
			ans = append(ans, tmp)

			return
		}
		if idx >= len(candidates) || target < 0 {
			return
		}
		for i := idx; i < len(candidates); i++ {
			if i > idx && candidates[i] == candidates[i-1] {
				// 如果第一个已经选用过了，那之后就跳过相同的值
				continue
			}
			*path = append(*path, candidates[i])
			// i不可重用，所以必须+1
			backtrack(target-candidates[i], i+1, path)
			*path = (*path)[:len(*path)-1]
		}
	}
	backtrack(target, 0, &[]int{})
	return ans
}

// @lc code=end
