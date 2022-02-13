/*
 * @lc app=leetcode.cn id=1011 lang=golang
 *
 * [1011] 在 D 天内送达包裹的能力
 */
package main

import "fmt"

// @lc code=start
func shipWithinDays(weights []int, days int) int {
	minCap := weights[0]
	maxCap := 0
	for _, weight := range weights {
		maxCap += weight
		minCap = max(minCap, weight)
	}
	for minCap <= maxCap {
		mid := (minCap + maxCap) / 2
		fmt.Println(mid)
		d := carry(weights, mid)
		if d < days {
			// 寻找左边界
			// 条件为l<r时，不能使用max = mid-1
			// 条件为l<=r时，两种收缩方式都用mid-1 or mid+1
			maxCap = mid - 1
		} else if d > days {
			minCap = mid + 1
		} else {
			maxCap = mid - 1
		}
	}
	return minCap
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func carry(weights []int, cap int) int {
	count := 1
	curCap := cap
	for _, weight := range weights {
		if curCap >= weight {
			curCap -= weight
		} else {
			count++
			curCap = cap - weight
		}
	}
	//fmt.Printf("cap %d day %d\n", cap, count)
	return count
}

// @lc code=end
