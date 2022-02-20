package main

import "fmt"

// A = [
//   [ 1, 0, 0],
//   [-1, 0, 3]
// ]

// B = [
//   [ 7, 0, 0 ],
//   [ 0, 0, 0 ],
//   [ 0, 0, 1 ]
// ]

//      |  1 0 0 |   | 7 0 0 |   |  7 0 0 |
// AB = | -1 0 3 | x | 0 0 0 | = | -7 0 3 |
//                   | 0 0 1 |
func multiply(a, b [][]int) [][]int {
	ans := make([][]int, len(a))
	for i := range ans {
		ans[i] = make([]int, len(b[0]))
	}
	for rowA := 0; rowA < len(a); rowA++ {
		for colB := 0; colB < len(b[0]); colB++ {
			for i := range a[rowA] {
				if a[rowA][i] != 0 && b[i][colB] != 0 {
					ans[rowA][colB] += a[rowA][i] * b[i][colB]
				}
			}
		}
	}

	return ans
}

func main() {
	a := [][]int{{1, 0, 0}, {-1, 0, 3}}
	b := [][]int{{7, 0, 0}, {0, 0, 0}, {0, 0, 1}}
	ans := multiply(a, b)
	for _, row := range ans {
		fmt.Printf("%v\n", row)
	}
}
