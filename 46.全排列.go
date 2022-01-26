/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 */
package main

// @lc code=start
func permute(nums []int) [][]int {
	result := new([][]int)
	cur := []int{}
	permuteCore(nums, result, cur)
	return *result
}

func permuteCore(nums []int, result *[][]int, cur []int) {
	if len(nums) == 1 {
		cur = append(cur, nums[0])
		*result = append(*result, cur)
		return
	}

	for i := 0; i < len(nums); i++ {
		newNums := []int{}
		newNums = append(newNums, nums[0:i]...)
		newNums = append(newNums, nums[i+1:]...)
		copyCur := []int{}
		copyCur = append(copyCur, cur...)
		copyCur = append(copyCur, nums[i])
		permuteCore(newNums, result, copyCur)
	}
}

// @lc code=end
