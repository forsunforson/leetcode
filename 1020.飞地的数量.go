/*
 * @lc app=leetcode.cn id=1020 lang=golang
 *
 * [1020] 飞地的数量
 */

// @lc code=start
func numEnclaves(grid [][]int) int {
	// union-find
	// 先将边界的陆地标注出来，然后寻找联通的单元
	// 找出非联通的陆地单元数量
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i != 0 && i != m-1 && j != 0 && j != n-1 {
				// 跳过非边界单元
				continue
			}
			isBoard(&grid, i, j)
		}
	}
	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				count++
			}
		}
	}
	return count
}

func isBoard(grid *[][]int, x, y int) {
	// 标注联通边界的陆地
	m, n := len(*grid), len((*grid)[0])
	if x < 0 || x >= m || y < 0 || y >= n {
		return
	}
	if (*grid)[x][y] == 1 {
		(*grid)[x][y] = -1
		isBoard(grid, x+1, y)
		isBoard(grid, x-1, y)
		isBoard(grid, x, y+1)
		isBoard(grid, x, y-1)
	}
}

// @lc code=end

