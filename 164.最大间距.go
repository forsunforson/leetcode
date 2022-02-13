/*
 * @lc app=leetcode.cn id=164 lang=golang
 *
 * [164] 最大间距
 */
package main

import "fmt"

// @lc code=start
func maximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	// 桶排序，间隔最大不会出现在桶内，只需比较桶间
	maxNum, minNum := nums[0], nums[0]
	for _, num := range nums {
		maxNum = max(maxNum, num)
		minNum = min(minNum, num)
	}
	// n+1个桶，每个间隔(max-min)/n
	meanGap := max(1, (maxNum-minNum)/len(nums))
	fmt.Printf("mean gap: %d", meanGap)
	buckets := make([][]int, (maxNum-minNum)/meanGap+1)
	for i := 0; i < len(buckets); i++ {
		buckets[i] = []int{}
	}
	for _, num := range nums {
		buckets[(num-minNum)/meanGap] = append(buckets[(num-minNum)/meanGap], num)
	}
	maxGap := 0
	lastMax := minNum
	for i := 0; i < len(buckets); i++ {
		if len(buckets[i]) == 0 {
			continue
		}
		curMin, curMax := buckets[i][0], buckets[i][0]
		for _, num := range buckets[i] {
			curMin = min(curMin, num)
			curMax = max(curMax, num)
		}
		if i != 0 {
			maxGap = max(maxGap, curMin-lastMax)
		}
		lastMax = curMax
	}
	return maxGap
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

// @lc code=end
