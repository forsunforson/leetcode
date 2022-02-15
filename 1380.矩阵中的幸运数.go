/*
 * @lc app=leetcode.cn id=1380 lang=golang
 *
 * [1380] 矩阵中的幸运数
 */
package main

// @lc code=start
func luckyNumbers(matrix [][]int) []int {
	ans := []int{}
	rowMin, colMax := make([]int, len(matrix)), make([]int, len(matrix[0]))

	for i, row := range matrix {
		rowMin[i] = row[0]
		for j, num := range row {
			rowMin[i] = min(rowMin[i], num)
			colMax[j] = max(colMax[j], num)
		}
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if rowMin[i] == colMax[j] {
				ans = append(ans, matrix[i][j])
			}
		}
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func notGood(matrix [][]int) []int {
	ans := []int{}
	rowMin := [][]int{}
	for i, row := range matrix {
		num := row[0]
		y := 0
		for j, col := range row {
			if num > col {
				num = col
				y = j
			}
		}
		rowMin = append(rowMin, []int{i, y})
	}
	for _, v := range rowMin {
		colMax := matrix[v[0]][v[1]]
		func() {
			for i := 0; i < len(matrix); i++ {
				if matrix[i][v[1]] > colMax {
					return
				}
			}
			ans = append(ans, colMax)
		}()
	}
	return ans
}

// @lc code=end
