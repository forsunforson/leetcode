/*
 * @lc app=leetcode.cn id=435 lang=golang
 *
 * [435] 无重叠区间
 */
package main

import (
	"math"
	"sort"
)

// @lc code=start
func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		// 不用排序左，只要右越小越好
		return intervals[i][1] < intervals[j][1]
	})
	//fmt.Printf("%v\n", intervals)
	// 可以视区间为[start, end)
	// len(intervals) - 最大无重叠区间
	uncover := 0
	curEnd := math.MinInt32
	for _, interval := range intervals {
		if interval[0] >= curEnd {
			uncover++
			curEnd = interval[1]
		}
	}
	return len(intervals) - uncover
}

// @lc code=end
