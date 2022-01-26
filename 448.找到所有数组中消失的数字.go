/*
 * @lc app=leetcode.cn id=448 lang=golang
 *
 * [448] 找到所有数组中消失的数字
 */
package main

// @lc code=start
func findDisappearedNumbers(nums []int) []int {
	// inexist := make(map[int]int)
	// for i := 1; i <= len(nums); i++ {
	// 	inexist[i] = 1
	// }
	// result := []int{}
	// for _, num := range nums {
	// 	inexist[num] = 0
	// }
	// for i := 1; i <= len(nums); i++ {
	// 	if inexist[i] != 0 {
	// 		result = append(result, i)
	// 	}
	// }
	return find2(nums)
}

func find2(nums []int) []int {
	n := len(nums)
	for _, num := range nums {
		num = (num - 1) % n
		nums[num] += n
	}
	ret := []int{}
	for i, num := range nums {
		if num <= n {
			ret = append(ret, i+1)
		}
	}
	return ret
}

// @lc code=end
