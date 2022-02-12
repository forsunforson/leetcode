/*
 * @lc app=leetcode.cn id=875 lang=golang
 *
 * [875] 爱吃香蕉的珂珂
 */

// @lc code=start
func minEatingSpeed(piles []int, h int) int {
	maxK := 0
	for _, pile := range piles {
		maxK += pile
	}
	minK := 1
	for maxK > minK {
		midK := (maxK + minK) / 2
		fmt.Printf("maxK: %d minK: %d\n", maxK, minK)
		if canFinish(h, midK, piles) {
			maxK = midK
		} else {
			minK = midK + 1
		}
	}
	return minK
}

func canFinish(h, k int, piles []int) bool {
	t := 0
	for _, pile := range piles {
		t += (pile-1)/k + 1
	}
	fmt.Printf("k: %d t: %d\n", k, t)
	return t <= h
}

// @lc code=end

