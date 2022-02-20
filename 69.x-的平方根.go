/*
 * @lc app=leetcode.cn id=69 lang=golang
 *
 * [69] x 的平方根
 */
package main

import "math"

// @lc code=start
func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	pre := 1
	for i := 1; i <= x/2; i++ {
		if i*i <= x {
			pre = i
		} else {
			break
		}
	}
	return pre
}

func newton(x int) int {
	if x == 0 {
		return 0
	}
	C, x0 := float64(x), float64(x)
	for {
		xi := 0.5 * (x0 + C/x0)
		if math.Abs(x0-xi) < 1e-7 {
			break
		}
		x0 = xi
	}
	return int(x0)
}

// @lc code=end
