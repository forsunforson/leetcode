/*
 * @lc app=leetcode.cn id=75 lang=golang
 *
 * [75] 颜色分类
 */
package main

// @lc code=start
func sortColors(nums []int) {
	p0, p1 := 0, 0
	for i, num := range nums {
		if num == 0 {
			nums[p0], nums[i] = nums[i], nums[p0]
			if p0 < p1 {
				nums[i], nums[p1] = nums[p1], nums[i]
			}
			p0++
			p1++
		} else if num == 1 {
			nums[p1], nums[i] = nums[i], nums[p1]
			p1++
		}
	}
}

// @lc code=end
