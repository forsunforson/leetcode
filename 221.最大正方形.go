/*
 * @lc app=leetcode.cn id=221 lang=golang
 *
 * [221] 最大正方形
 */
package main

// @lc code=start
func maximalSquare(matrix [][]byte) int {
	dp := make([][]int, len(matrix))
	for i := range dp {
		dp[i] = make([]int, len(matrix[0]))
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == '0' {
				dp[i][j] = 0
			} else {
				if j > 0 {
					dp[i][j] = dp[i][j-1] + 1
				} else {
					dp[i][j] = 1
				}
			}
		}
	}

	max := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			curMin := dp[i][j]
			for k := i; k < len(matrix); k++ {
				if dp[k][j] < curMin {
					curMin = dp[k][j]
				}
				length := min(k-i+1, curMin)
				curAread := length * length
				if max < curAread {
					max = curAread
				}
			}
		}
	}
	return max
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// @lc code=end
