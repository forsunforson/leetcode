/*
 * @lc app=leetcode.cn id=48 lang=golang
 *
 * [48] 旋转图像
 */
package main

// @lc code=start
func rotate(matrix [][]int) {
	for i := 0; i < len(matrix)/2; i++ {
		for j := i; j < len(matrix)-i-1; j++ {
			rotateOnePix(matrix, i, j)
		}
	}
}

// (0,0) -> (0,2) (0,2) -> (2,2) (2,2) -> (2,0)
// (0,1) -> (1,2) (1,2) -> (2,1)
func rotateOnePix(matrix [][]int, x, y int) {
	tmpPix := matrix[x][y]
	n := len(matrix)
	newX, newY := x, y
	for i := 0; i < 4; i++ {
		newX, newY = getNexPos(n, newX, newY)
		pix := matrix[newX][newY]
		matrix[newX][newY] = tmpPix
		tmpPix = pix
	}
}

func getNexPos(n, x, y int) (int, int) {
	return y, n - x - 1
}

// @lc code=end
