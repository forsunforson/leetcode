/*
 * @lc app=leetcode.cn id=200 lang=golang
 *
 * [200] 岛屿数量
 */
package main

import "fmt"

// @lc code=start
func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	// 初始化map，记录是否已访问过
	rowLen := len(grid)
	colLen := len(grid[0])
	gridMap := make([][]bool, rowLen)
	for i := 0; i < rowLen; i++ {
		gridMap[i] = make([]bool, len(grid[0]))
	}
	counter := 0
	// 遍历grid，如果没有访问过则递归访问，增加计数
	for i := 0; i < rowLen; i++ {
		for j := 0; j < colLen; j++ {
			if grid[i][j] == byte('0') {
				gridMap[i][j] = true
				continue
			}
			if gridMap[i][j] {
				continue
			}
			visit(grid, &gridMap, i, j)
			counter++
		}
	}
	return counter
}

func visit(grid [][]byte, gridMap *[][]bool, i, j int) {
	rowLen := len(grid)
	colLen := len(grid[0])
	if i < 0 || i >= rowLen || j < 0 || j >= colLen {
		return
	}
	if (*gridMap)[i][j] {
		return
	}
	if grid[i][j] == byte('0') {
		return
	}
	fmt.Printf("visit [%d][%d]\n", i, j)
	(*gridMap)[i][j] = true
	visit(grid, gridMap, i+1, j)
	visit(grid, gridMap, i, j+1)
	visit(grid, gridMap, i-1, j)
	visit(grid, gridMap, i, j-1)
}

// @lc code=end
