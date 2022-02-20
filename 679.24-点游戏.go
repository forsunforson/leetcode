/*
 * @lc app=leetcode.cn id=679 lang=golang
 *
 * [679] 24 点游戏
 */
package main

// @lc code=start
func judgePoint24(cards []int) bool {
	// 选出任意两个数做四则运算，暴力
	// np问题，回溯法
	nums := []float64{}
	for _, card := range cards {
		nums = append(nums, float64(card))
	}
	var backtrack func([]float64) bool
	backtrack = func(f []float64) bool {
		// 从数中任意选出2个，进行四则运算后放回
		if len(f) == 0 {
			return false
		}
		if len(f) == 1 {
			return isNear(f[0], 24)
		}
		size := len(f)
		for i := 0; i < size; i++ {
			for j := 0; j < size; j++ {
				if i == j {
					continue
				}
				newF := []float64{}
				// 把没选到的加到newF里
				for k := 0; k < size; k++ {
					if k == i || k == j {
						continue
					}
					newF = append(newF, f[k])
				}
				// 对i j执行四则运算
				for k := 0; k < 4; k++ {
					switch k {
					case 0:
						newF = append(newF, f[i]+f[j])
					case 1:
						newF = append(newF, f[i]-f[j])
					case 2:
						newF = append(newF, f[i]*f[j])
					case 3:
						if !isNear(f[j], 0) {
							newF = append(newF, f[i]/f[j])
						} else {
							continue
						}
					}
					if backtrack(newF) {
						return true
					}
					newF = newF[:len(newF)-1]
				}
			}
		}
		return false
	}
	return backtrack(nums)
}

func isNear(i float64, j int) bool {
	if i-float64(j) <= 1e-6 && i-float64(j) >= -1e-6 {
		return true
	}
	return false
}

// @lc code=end
