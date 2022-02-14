/*
 * @lc app=leetcode.cn id=773 lang=golang
 *
 * [773] 滑动谜题
 */
package main

import "fmt"

// @lc code=start
type step struct {
	board string
	count int
}

var neighbors = [6][]int{{1, 3}, {0, 2, 4}, {1, 5}, {0, 4}, {1, 3, 5}, {2, 4}}

func slidingPuzzle(board [][]int) int {
	// 找最优解，就是用BFS
	// 这道题永远是用0和其他位置交换
	queue := []step{}
	target := "123450"
	visited := map[string]bool{}
	str := ""
	for _, row := range board {
		for _, col := range row {
			str += fmt.Sprint(col)
		}
	}
	if target == str {
		return 0
	}
	queue = append(queue, step{
		board: str,
		count: 0,
	})
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		for _, s := range cur.generate() {
			if !visited[s.board] {
				if target == s.board {
					return s.count
				}
				visited[s.board] = true
				queue = append(queue, s)
			}
		}
	}
	return -1
}

func (s step) generate() []step {
	steps := []step{}
	idx := 0
	for s.board[idx] != '0' {
		idx++
	}
	b := []byte(s.board)
	for _, neighbor := range neighbors[idx] {
		b[idx], b[neighbor] = b[neighbor], b[idx]
		steps = append(steps, step{string(b), s.count + 1})
		b[idx], b[neighbor] = b[neighbor], b[idx]
	}
	return steps
}

// @lc code=end
