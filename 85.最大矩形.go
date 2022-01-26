/*
 * @lc app=leetcode.cn id=85 lang=golang
 *
 * [85] 最大矩形
 */
package main

// @lc code=start
func maximalRectangle(matrix [][]byte) int {
	// 遍历一次矩阵，把它每一行化作heights[]int
	maxSquare := 0
	var lastRow []int
	for i := 0; i < len(matrix); i++ {
		//fmt.Printf("%v\n", matrix[i])
		row := make([]int, len(matrix[0]))
		for j := 0; j < len(row); j++ {
			if i == 0 {
				if matrix[i][j] == 48 {
					row[j] = 0
				} else {
					row[j] = 1
				}
			} else {
				if matrix[i][j] == 48 {
					row[j] = 0
				} else {
					row[j] = lastRow[j] + 1
				}
			}
		}
		square := getMaxSquare(row)
		if square > maxSquare {
			maxSquare = square
		}
		lastRow = row
	}
	return maxSquare
}

func getMaxSquare(heights []int) int {
	maxSquare := 0
	left := make([]int, len(heights))
	right := make([]int, len(heights))
	leftStack := new([]int)
	rightStack := new([]int)
	for i := 0; i < len(heights); i++ {
		left[i] = pushStack(leftStack, heights, i, false)
		right[len(heights)-1-i] = pushStack(rightStack, heights, len(heights)-i-1, true)
	}
	for i := 0; i < len(heights); i++ {
		square := (right[i] - left[i] - 1) * heights[i]
		if maxSquare < square {
			maxSquare = square
		}
	}
	return maxSquare
}

func pushStack(stack *[]int, heights []int, idx int, reversed bool) int {
	for len(*stack) > 0 {
		if heights[(*stack)[len(*stack)-1]] >= heights[idx] {
			*stack = (*stack)[:len(*stack)-1]
		} else {
			break
		}
	}
	*stack = append(*stack, idx)
	if len(*stack) == 1 {
		if reversed {
			return len(heights)
		} else {
			return -1
		}
	}
	return (*stack)[len(*stack)-2]
}

// @lc code=end
