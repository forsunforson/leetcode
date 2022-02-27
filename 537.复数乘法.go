/*
 * @lc app=leetcode.cn id=537 lang=golang
 *
 * [537] 复数乘法
 */
package main

import (
	"fmt"
	"strconv"
	"strings"
)

// @lc code=start
func complexNumberMultiply(num1 string, num2 string) string {
	n1, n2 := parseComp(num1), parseComp(num2)
	res1 := n1[0]*n2[0] - n1[1]*n2[1]
	res2 := n1[0]*n2[1] + n1[1]*n2[0]
	return fmt.Sprintf("%d+%di", res1, res2)
}
func parseComp(num string) []int {
	s := strings.Split(num, "+")

	n1, _ := strconv.Atoi(s[0])
	n2, _ := strconv.Atoi(s[1][:len(s[1])-1])
	return []int{n1, n2}
}

// @lc code=end
