/*
 * @lc app=leetcode.cn id=134 lang=golang
 *
 * [134] 加油站
 */
package main

// @lc code=start
func canCompleteCircuit(gas []int, cost []int) int {
	rest := make([]int, len(gas))
	minIdx := 0
	for i := range gas {
		if i > 0 {
			rest[i] = rest[i-1] + gas[i] - cost[i]
		} else {
			rest[i] = gas[i] - cost[i]
		}
		if rest[minIdx] > rest[i] {
			minIdx = i
		}
	}
	// 从最小的下一站开始
	sum := 0
	for i := 1; i <= len(gas); i++ {
		start := minIdx + i
		if start >= len(gas) {
			start -= len(gas)
		}
		sum += gas[start] - cost[start]
		if sum < 0 {
			return -1
		}
	}
	if minIdx+1 >= len(gas) {
		minIdx -= len(gas)
	}
	return minIdx + 1
}

// @lc code=end
