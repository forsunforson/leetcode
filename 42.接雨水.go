/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 */
package main

// @lc code=start
func trap(height []int) int {
	// 双指针
	ans := 0
	l, r := 0, len(height)-1
	lMax, rMax := 0, 0
	for {
		if l >= r {
			break
		}
		if height[l] < height[r] {
			if lMax < height[l] {
				lMax = height[l]
			}
			ans += lMax - height[l]
			l++
		} else {
			if rMax < height[r] {
				rMax = height[r]
			}
			ans += rMax - height[r]
			r--
		}
	}
	return ans
}

func trap2(height []int) int {
	// 动态规划
	ans := 0
	leftStack := make([]int, len(height))
	rightStack := make([]int, len(height))
	leftMax, rightMax := 0, 0
	for i := 0; i < len(height); i++ {
		if i == 0 {
			leftMax = height[i]
			rightMax = height[len(height)-i-1]
			leftStack[i] = leftMax
			rightStack[len(height)-i-1] = rightMax
			continue
		}
		if height[i] > leftMax {
			leftMax = height[i]
		}
		leftStack[i] = leftMax
		if height[len(height)-i-1] > rightMax {
			rightMax = height[len(height)-i-1]
		}
		rightStack[len(height)-i-1] = rightMax
	}
	for i := 0; i < len(height); i++ {
		h := leftStack[i]
		if rightStack[i] < h {
			h = rightStack[i]
		}
		ans += h - height[i]
	}
	return ans
}

func trap1(height []int) int {
	// 单调栈
	stack := make([]int, 0)
	ans := 0
	for i := 0; i < len(height); i++ {
		for {
			// 循环更新比当前值小的值
			if len(stack) == 0 {
				break
			}
			top := stack[len(stack)-1]
			if height[top] > height[i] {
				break
			}
			stack = stack[:len(stack)-1]
			h := height[i]
			if len(stack) == 0 {
				break
			}
			left := stack[len(stack)-1]

			if height[left] < h {
				h = height[left]
			}
			ans += (i - left - 1) * (h - height[top])
		}
		stack = append(stack, i)
	}
	return ans
}

// @lc code=end
