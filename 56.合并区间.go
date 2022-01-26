/*
 * @lc app=leetcode.cn id=56 lang=golang
 *
 * [56] 合并区间
 */
package main

import (
	"fmt"
	"sort"
)

// @lc code=start

type sortInt [][]int

func (s sortInt) Len() int {
	return len(s)
}
func (s sortInt) Swap(i, j int) {
	tmp := s[i]
	s[i] = s[j]
	s[j] = tmp
}
func (s sortInt) Less(i, j int) bool {
	if s[i][0] != s[j][0] {
		return s[i][0] < s[j][0]
	}
	return s[i][1] < s[j][1]
}
func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}
	sortedIntervals := sortInt(intervals)
	sort.Sort(sortedIntervals)
	fmt.Printf("%v\n", sortedIntervals)
	ret := make([][]int, 0)
	for i := 0; i < len(intervals); {
		curStart, curEnd := sortedIntervals[i][0], sortedIntervals[i][1]
		nextStart, nextEnd := sortedIntervals[i][0], sortedIntervals[i][1]
		j := i + 1
		for j < len(intervals) {
			nextStart, nextEnd = sortedIntervals[j][0], sortedIntervals[j][1]
			if curEnd < nextStart {
				break
			}
			if curEnd < nextEnd {
				curEnd = nextEnd
			}
			j++
		}
		ret = append(ret, []int{curStart, curEnd})
		i = j
	}
	return ret
}

// @lc code=end
