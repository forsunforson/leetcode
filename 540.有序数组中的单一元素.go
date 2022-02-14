/*
 * @lc app=leetcode.cn id=540 lang=golang
 *
 * [540] 有序数组中的单一元素
 */
package main

// @lc code=start
func singleNonDuplicate(nums []int) int {
	// O(log n)
	// 二分查找
	l, r := 0, len(nums)-1
	// 一般使用[l, r)区间，所以l=mid+1, r=mid
	for l < r {
		mid := (l + r) / 2
		// 假如不存在落单数
		// mid偶数，比较mid与mid-1
		// mid奇数，比较mid与mid+1
		next := mid ^ 1
		if nums[next] == nums[mid] {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return nums[l]
}

// O(n)
func s1(nums []int) int {
	ans := 0
	for _, num := range nums {
		ans ^= num
	}
	return ans
}

// @lc code=end
