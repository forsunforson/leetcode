/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 */
package main

// @lc code=start
func subsets(nums []int) [][]int {
	results := make([][]int, 0)
	result := []int{}
	subsetsHelper(0, nums, result, &results)
	return results
}

func subsetsHelper(idx int, nums, result []int, results *[][]int) {
	if idx == len(nums) {
		*results = append(*results, append([]int{}, result...))
		return
	}
	// 加
	result = append(result, nums[idx])
	subsetsHelper(idx+1, nums, result, results)
	// 不加
	result = result[:len(result)-1]
	subsetsHelper(idx+1, nums, result, results)
}

// @lc code=end
