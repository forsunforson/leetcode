package main

import "fmt"

// 暴力破解是一种
// 但一般都不用，用二分法

// 1 <= A.length <= 50000
// 1 <= A[i] <= 1e7
// 1 <= K <= 1e8
func findMissingKth(nums []int, k int) int {
	l, r := 0, len(nums)-1
	for l < r-1 {
		mid := (r-l)/2 + l
		numCountL := nums[mid] - nums[l] - mid + l
		if numCountL < k {
			// 缺失数在右区间
			l = mid
			k -= numCountL
		} else {
			// 缺失数在左区间
			r = mid
		}

	}
	if k > nums[r]-nums[l] {
		return nums[l] + k + 1
	}
	return nums[l] + k
}

func main() {
	fmt.Printf("%v\n", findMissingKth([]int{1, 2, 4}, 3))
	fmt.Printf("%v\n", findMissingKth([]int{4, 7, 9, 10}, 3))
	fmt.Printf("%v\n", findMissingKth([]int{4, 7, 9, 10}, 1))
}
