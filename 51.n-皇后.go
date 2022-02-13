/*
 * @lc app=leetcode.cn id=51 lang=golang
 *
 * [51] N 皇后
 */
package main

import "strings"

// @lc code=start
func solveNQueens(n int) [][]string {
	// 问题的解决思路比较简单
	// 使用回溯法回溯所有的可能性，选择其中能够符合规则的
	// 本题的难点在于如何简单的判断皇后是否能相互攻击
	// 在放置时，就是按不同行进行放置的，所以考察列就可以
	// 对于斜向的，是否在一条斜线上，可以由x+y或者x-y是否相等来进行判断
	ans := [][]string{}
	column := make(map[int]bool)
	diag1 := make(map[int]bool)
	diag2 := make(map[int]bool)
	queensCore(&ans, n, 0, [][]string{}, column, diag1, diag2)
	return ans
}

func queensCore(ans *[][]string, n, idx int, cur [][]string, column, diag1, diag2 map[int]bool) {
	if idx == n {
		output := []string{}
		for _, str := range cur {
			output = append(output, strings.Join(str, ""))
		}
		*ans = append(*ans, output)
		return
	}

	for i := 0; i < n; i++ {
		if column[i] || diag1[idx+i] || diag2[idx-i] {
			continue
		}
		curStr := []string{}
		for j := 0; j < n; j++ {
			curStr = append(curStr, ".")
		}
		curStr[i] = "Q"
		cur = append(cur, curStr)
		column[i] = true
		diag1[idx+i] = true
		diag2[idx-i] = true
		queensCore(ans, n, idx+1, cur, column, diag1, diag2)
		cur = cur[:len(cur)-1]
		column[i] = false
		diag1[idx+i] = false
		diag2[idx-i] = false
	}
}

// @lc code=end
