/*
 * @lc app=leetcode.cn id=1095 lang=golang
 *
 * [1095] 山脉数组中查找目标值
 */

// @lc code=start
/**
 * // This is the MountainArray's API interface.
 * // You should not implement it, or speculate about its implementation
 * type MountainArray struct {
 * }
 *
 * func (this *MountainArray) get(index int) int {}
 * func (this *MountainArray) length() int {}
 */

func findInMountainArray(target int, mountainArr *MountainArray) int {
	// 二分法，mid=(l+r)/2
	// 假设左半有序，arr[mid-1]<arr[mid]，从左半中二分查找，找到返回，找不到返回递归右半
	// 假设右半有序，递归左半
	l, r := 0, mountainArr.length()-1
	return find(target, l, r, mountainArr)
}

func find(target, l, r int, arr *MountainArray) int {
	if l == r {
		if arr.get(l) == target {
			return l
		}
		return -1
	}
	mid := (l + r) / 2
	if mid == 0 {
		if arr.get(mid) == target {
			return 0
		} else {
			return find(target, 1, r, arr)
		}
	}
	if arr.get(mid-1) < arr.get(mid) {
		// 左半递增
		if arr.get(mid) < target {
			return find(target, mid+1, r, arr)
		}
		// 先从左半找，找不到再找右半
		start, end := l, mid
		idx := biSearch(target, start, end, arr)
		if idx != -1 {
			return idx
		} else {
			return find(target, mid+1, r, arr)
		}
	} else {
		// 右半递减
		idx := find(target, l, mid, arr)
		if idx != -1 {
			return idx
		}
		// 左半找不到，右半二分
		start, end := mid+1, r
		for start <= end {
			newMid := (start + end) / 2
			n := arr.get(newMid)
			if n == target {
				return newMid
			}
			if n < target {
				end = newMid - 1
			} else {
				start = newMid + 1
			}
		}
		return -1
	}
}

func biSearch(target, l, r int, arr *MountainArray) int {
	for l <= r {
		mid := (l + r) / 2
		n := arr.get(mid)
		if n == target {
			return mid
		}
		if n < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return -1
}

// @lc code=end

