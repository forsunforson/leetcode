package main

import "fmt"

// Given matrix = [
//   [3, 0, 1, 4, 2],
//   [5, 6, 3, 2, 1],
//   [1, 2, 0, 1, 5],
//   [4, 1, 0, 1, 7],
//   [1, 0, 3, 0, 5]
// ]

// sumRegion(2, 1, 4, 3) -> 8
// update(3, 2, 2)
// sumRegion(2, 1, 4, 3) -> 10

type Matrix struct {
	m    [][]int
	cols [][]int
}

// 使用前缀树
func NewMatrix(matrix [][]int) Matrix {
	m := Matrix{m: matrix}
	cols := make([][]int, len(matrix))
	for i := range cols {
		cols[i] = make([]int, len(matrix[0]))
		sum := 0
		for j := range cols[i] {
			sum += matrix[i][j]
			cols[i][j] = sum
		}
	}
	m.cols = cols
	return m
}

func (m *Matrix) SumRegion(upX, upY, botX, botY int) int {
	sum := 0
	for x := upX; x <= botX; x++ {
		sum += m.cols[x][botY]
		if upY > 0 {
			sum -= m.cols[x][upY-1]
		}
	}
	return sum
}

func (m *Matrix) Update(x, y, val int) {
	old := m.m[x][y]
	m.m[x][y] = val
	for i := y; i < len(m.m[0]); i++ {
		m.cols[x][i] = m.cols[x][i] - old + val
	}
}

func main() {
	matrix := [][]int{
		{3, 0, 1, 4, 2},
		{5, 6, 3, 2, 1},
		{1, 2, 0, 1, 5},
		{4, 1, 0, 1, 7},
		{1, 0, 3, 0, 5},
	}
	m := NewMatrix(matrix)
	fmt.Println(m.SumRegion(2, 1, 4, 3))
	m.Update(3, 2, 2)
	fmt.Println(m.SumRegion(2, 1, 4, 3))
}
