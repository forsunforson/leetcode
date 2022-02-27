/*
 * @lc app=leetcode.cn id=1706 lang=golang
 *
 * [1706] 球会落何处
 */
package main

import "fmt"

// @lc code=start
func findBall(grid [][]int) []int {
	// 当前坐标(x,y)，值是1，检查(x,y+1)是否是-1,是则卡住；否则进入(x+1,y+1)
	ans := make([]int, len(grid[0]))
	for idx := range grid[0] {
		end := -1

		i, j := idx, 0
		for {
			fmt.Print("(", j, " ", i, ")")
			if j >= len(grid) {
				end = i
				break
			}
			if grid[j][i] == 1 {
				if i == len(grid[j])-1 || grid[j][i+1] == -1 {
					end = -1
					break
				}
				j++
				i++
			} else {
				if i == 0 || grid[j][i-1] == 1 {
					end = -1
					break
				}
				j++
				i--
			}
		}
		ans[idx] = end
		fmt.Println()
	}
	return ans
}

// @lc code=end
