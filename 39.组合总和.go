/*
 * @lc app=leetcode.cn id=39 lang=golang
 *
 * [39] 组合总和
 */
package main

import (
	"fmt"
	"sort"
)

// @lc code=start
func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	ret := new([][]int)
	cur := make([]int, 0)
	combinationCore(candidates, target, ret, cur, 0)
	return *ret
}

func combinationCore(candidates []int, target int, ret *[][]int, cur []int, i int) {
	//fmt.Printf("input:%v ", cur)
	if target == 0 {
		*ret = append(*ret, cur)
		fmt.Printf("ret:%v\n", cur)
		return
	}
	if target < 0 {
		return
	}
	for j := i; j < len(candidates); j++ {
		curNum := candidates[j]
		//fmt.Printf("curNum:%d ", curNum)
		if curNum <= target {
			copyCur := []int{}
			copyCur = append(copyCur, cur...)
			copyCur = append(copyCur, curNum)
			//fmt.Printf("output:%v\n", copyCur)
			combinationCore(candidates, target-curNum, ret, copyCur, j)
		} else {
			return
		}
	}
}

// @lc code=end
