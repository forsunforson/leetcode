/*
 * @lc app=leetcode.cn id=73 lang=golang
 *
 * [73] 矩阵置零
 */
package main

// @lc code=start
func setZeroes(matrix [][]int) {
	// m n范围[1,200]，可以用两个[25]byte来记录0所在的行、列号
	rowBit := make([]byte, 25)
	colBit := make([]byte, 25)
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == 0 {
				rowBit[i/8] |= 1 << (i % 8)
				colBit[j/8] |= 1 << (j % 8)
			}
		}
	}
	//fmt.Printf("%v %v\n", rowBit, colBit)
	for i, bit := range rowBit {
		count := 0
		for bit > 0 {
			if bit&1 != 0 {
				row := i*8 + count
				//fmt.Printf("row:%d\n", row)
				for j := range matrix[row] {
					matrix[row][j] = 0
				}
			}
			bit >>= 1
			count++
		}
	}
	for i, bit := range colBit {
		count := 0
		for bit > 0 {
			if bit&1 != 0 {
				col := i*8 + count
				//fmt.Printf("col:%d\n", col)
				for j := range matrix {
					matrix[j][col] = 0
				}
			}
			bit >>= 1
			count++
		}
	}
}

// @lc code=end
