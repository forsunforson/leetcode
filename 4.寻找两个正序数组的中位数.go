/*
 * @lc app=leetcode.cn id=4 lang=golang
 *
 * [4] 寻找两个正序数组的中位数
 */
package main

import "fmt"

// @lc code=start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n, m := len(nums1), len(nums2)
	k := 0
	isEven := (n+m)%2 == 1
	if isEven {
		k = (n+m)/2 + 1
	} else {
		k = (n + m) / 2
	}

	if isEven {
		return float64(findK(nums1, nums2, k))
	}
	return (float64(findK(nums1, nums2, k)) + float64(findK(nums1, nums2, k+1))) / 2.0
}

func findK(nums1, nums2 []int, k int) int {
	idx1, idx2 := 0, 0

	for {
		if len(nums1) == idx1 {
			return nums2[idx2+k-1]
		}
		if len(nums2) == idx2 {
			return nums1[idx1+k-1]
		}
		if k == 1 {
			if nums1[idx1] < nums2[idx2] {
				return nums1[idx1]
			} else {
				return nums2[idx2]
			}
		}
		newIdx1 := len(nums1) - 1
		if newIdx1 > idx1+k/2-1 {
			newIdx1 = idx1 + k/2 - 1
		}
		newIdx2 := len(nums2) - 1
		if newIdx2 > idx2+k/2-1 {
			newIdx2 = idx2 + k/2 - 1
		}
		fmt.Printf("idx1 %d newIdx1 %d idx2 %d newIdx2 %d\n", idx1, newIdx1, idx2, newIdx2)
		if nums1[newIdx1] < nums2[newIdx2] {
			k -= (newIdx1 - idx1 + 1)
			idx1 = newIdx1 + 1
		} else {
			k -= (newIdx2 - idx2 + 1)
			idx2 = newIdx2 + 1
		}
	}
}

func solution1(nums1 []int, nums2 []int) float64 {
	// o(n+m)
	totalLen := len(nums1) + len(nums2)
	isEven := false
	mid := totalLen / 2
	if totalLen%2 == 1 {
		isEven = true
	}
	pre := 0
	cur := 0
	idx1, idx2 := 0, 0
	for i := 0; i <= mid; i++ {
		pre = cur
		if idx1 >= len(nums1) {
			cur = nums2[idx2]
			idx2++
			continue
		}
		if idx2 >= len(nums2) {
			cur = nums1[idx1]
			idx1++
			continue
		}
		if nums1[idx1] < nums2[idx2] {
			cur = nums1[idx1]
			idx1++
		} else {
			cur = nums2[idx2]
			idx2++
		}
	}
	if isEven {
		return float64(cur)
	} else {
		return (float64(pre) + float64(cur)) / 2.0
	}
}

// @lc code=end
