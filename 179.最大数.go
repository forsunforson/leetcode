/*
 * @lc app=leetcode.cn id=179 lang=golang
 *
 * [179] 最大数
 */
package main

import (
	"sort"
	"strconv"
)

// @lc code=start

type StrInt []int

func (a StrInt) Len() int      { return len(a) }
func (a StrInt) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a StrInt) Less(i, j int) bool {
	x := strconv.Itoa(a[i]) + strconv.Itoa(a[j])
	y := strconv.Itoa(a[j]) + strconv.Itoa(a[i])
	return x > y
}
func largestNumber(nums []int) string {
	ans := ""
	strInt := StrInt(nums)
	sort.Sort(strInt)
	if strInt[0] == 0 {
		return "0"
	}
	for _, num := range strInt {
		ans += strconv.Itoa(num)
	}

	return ans
}

// @lc code=end
