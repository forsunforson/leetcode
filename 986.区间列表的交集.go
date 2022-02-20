/*
 * @lc app=leetcode.cn id=986 lang=golang
 *
 * [986] 区间列表的交集
 */
package main

// @lc code=start
func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	if len(firstList) == 0 || len(secondList) == 0 {
		return [][]int{}
	}
	ans := [][]int{}
	idx1, idx2 := 0, 0
	for ; idx1 < len(firstList); idx1++ {
		start1, end1 := firstList[idx1][0], firstList[idx1][1]
		for idx2 < len(secondList) {
			// 有交集判断一下，没有交集idx2++
			start2, end2 := secondList[idx2][0], secondList[idx2][1]
			if start1 > end2 {
				idx2++
				continue
			}
			if start2 <= end1 {
				start := max(start1, start2)
				end := min(end1, end2)
				ans = append(ans, []int{start, end})
			}
			if end1 >= end2 {
				idx2++
			} else {
				break
			}
		}
	}
	return ans
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// @lc code=end
