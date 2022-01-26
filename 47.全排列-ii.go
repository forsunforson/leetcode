/*
 * @lc app=leetcode.cn id=47 lang=golang
 *
 * [47] 全排列 II
 */
package main

import "fmt"

// @lc code=start
type void struct{}

func permuteUnique(nums []int) [][]int {
	uniqueMap := make(map[string]interface{})
	result := new([][]int)
	cur := []int{}
	permuteUniqueCore(nums, cur, result, uniqueMap)
	return *result
}
func permuteUniqueCore(nums, cur []int, result *[][]int, uniqueMap map[string]interface{}) {
	if len(nums) == 1 {
		cur = append(cur, nums[0])
		uniqueKey := fmt.Sprintf("%v", cur)
		_, ok := uniqueMap[uniqueKey]
		if !ok {
			*result = append(*result, cur)
			uniqueMap[uniqueKey] = void{}
		}
		return
	}
	for i := 0; i < len(nums); i++ {
		newNums := []int{}
		newNums = append(newNums, nums[0:i]...)
		newNums = append(newNums, nums[i+1:]...)
		copyCur := []int{}
		copyCur = append(copyCur, cur...)
		copyCur = append(copyCur, nums[i])
		permuteUniqueCore(newNums, copyCur, result, uniqueMap)
	}

}

// @lc code=end
