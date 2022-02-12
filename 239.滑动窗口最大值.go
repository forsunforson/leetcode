/*
 * @lc app=leetcode.cn id=239 lang=golang
 *
 * [239] 滑动窗口最大值
 */

// @lc code=start
func maxSlidingWindow(nums []int, k int) []int {
	// 使用单调队列实现，入队列的数一定比队列内的要小
	// 随着滑动窗口，排出队头的idx
	// 队列内存储idx

	// 优化：把单调队列包装成对象
	if len(nums) == 0 || k == 0 {
		return []int{}
	}
	l, r := 0, 0
	queue := []int{}
	for r = 0; r < k; r++ {
		for {
			if len(queue) == 0 {
				break
			}
			lastIdx := queue[len(queue)-1]
			if nums[lastIdx] <= nums[r] {
				queue = queue[:len(queue)-1]
			} else {
				break
			}
		}

		queue = append(queue, r)
	}
	ans := []int{nums[queue[0]]}
	r = k - 1

	for {
		if l == queue[0] {
			queue = queue[1:]
		}
		l++
		r++
		if r >= len(nums) {
			break
		}
		for {
			if len(queue) == 0 {
				break
			}
			lastIdx := queue[len(queue)-1]
			if nums[lastIdx] <= nums[r] {
				queue = queue[:len(queue)-1]
			} else {
				break
			}
		}

		queue = append(queue, r)
		ans = append(ans, nums[queue[0]])
	}
	return ans
}

// @lc code=end

