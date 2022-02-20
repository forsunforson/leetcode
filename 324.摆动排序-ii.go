/*
 * @lc app=leetcode.cn id=324 lang=golang
 *
 * [324] 摆动排序 II
 */
package main

import "sort"

// @lc code=start
func wiggleSort(nums []int) {
	ans := append([]int{}, nums...)
	sort.Ints(ans)
	idx := (len(ans) + 1) / 2
	left := ans[:idx]
	right := ans[idx:]
	reverse(left)
	reverse(right)
	for i, num := range left {
		nums[i*2] = num
	}
	for i, num := range right {
		nums[i*2+1] = num
	}
}

func reverse(arr []int) {
	l, r := 0, len(arr)-1
	for l < r {
		arr[l], arr[r] = arr[r], arr[l]
		l++
		r--
	}
}

// @lc code=end
