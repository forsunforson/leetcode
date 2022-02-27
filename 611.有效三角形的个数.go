/*
 * @lc app=leetcode.cn id=611 lang=golang
 *
 * [611] 有效三角形的个数
 */
package main

import "sort"

// @lc code=start
func triangleNumber(nums []int) int {
	// 三条边 两边之和大于第三边 两边之差小于第三边
	sort.Ints(nums)
	count := 0
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			// 对于第三个数进行二分查找
			l, r := j+1, len(nums)-1
			k := j
			for l <= r {
				mid := (l + r) / 2
				if nums[mid] >= nums[i]+nums[j] {
					r = mid - 1
				} else {
					l = mid + 1
					k = mid
				}
			}
			count += k - j
		}
	}
	return count
}

// @lc code=end
func tooSlow(nums []int) int {
	// 三条边 两边之和大于第三边 两边之差小于第三边
	sort.Ints(nums)
	count := 0
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j] <= nums[k] {
					break
				}
				count++
			}
		}
	}
	return count
}
