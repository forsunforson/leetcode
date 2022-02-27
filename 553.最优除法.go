/*
 * @lc app=leetcode.cn id=553 lang=golang
 *
 * [553] 最优除法
 */
package main

import (
	"fmt"
	"strconv"
)

// @lc code=start
func optimalDivision(nums []int) string {
	// 除数越小越好，将数组分为a、b，a是第一个数字，后面为b。
	// 因为nums为正整数，所以连除不会变大
	if len(nums) == 1 {
		return fmt.Sprint(nums[0])
	}
	ans := ""
	substr := ""
	for i := 1; i < len(nums); i++ {
		substr += strconv.Itoa(nums[i])
		if i != len(nums)-1 {
			substr += "/"
		}
	}
	if len(nums) > 2 {
		substr = fmt.Sprintf("(%s)", substr)
	}
	ans += fmt.Sprintf("%d/%s", nums[0], substr)
	return ans
}

// @lc code=end
