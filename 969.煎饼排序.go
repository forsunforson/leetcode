/*
 * @lc app=leetcode.cn id=969 lang=golang
 *
 * [969] 煎饼排序
 */
package main

import "fmt"

// @lc code=start
func pancakeSort(arr []int) []int {
	ans := []int{}
	// 每次找到最大的饼，让它放到最下面
	last := len(arr)
	for !isSorted(arr) && last > 0 {
		maxNum := arr[0]
		maxIdx := 0
		for i := 0; i < last; i++ {
			if maxNum < arr[i] {
				maxIdx = i
				maxNum = arr[i]
			}
		}
		if maxIdx != 0 {
			ans = append(ans, maxIdx+1)
			arr = reverse(arr, maxIdx)
		}
		ans = append(ans, last)
		arr = reverse(arr, last-1)

		last--
		fmt.Printf("%v\n", arr)
	}
	return ans
}

func reverse(arr []int, idx int) []int {
	ret := make([]int, 0)
	for i := idx; i >= 0; i-- {
		ret = append(ret, arr[i])
	}
	for i := idx + 1; i < len(arr); i++ {
		ret = append(ret, arr[i])
	}
	return ret
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func isSorted(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i-1] > arr[i] {
			return false
		}
	}
	return true
}

// @lc code=end
