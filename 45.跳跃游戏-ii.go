/*
 * @lc app=leetcode.cn id=45 lang=golang
 *
 * [45] 跳跃游戏 II
 */

// @lc code=start
func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	// end 当前位置最远，furthest 当前步数最远
	end, furthest := 0, 0
	step := 0
	// 最后一个位置前就得结束
	for idx := 0; idx < len(nums)-1; idx++ {
		if furthest < idx+nums[idx] {
			furthest = idx + nums[idx]
		}
		if idx == end {
			step++
			end = furthest
		}

	}
	return step
}

// @lc code=end

