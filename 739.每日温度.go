/*
 * @lc app=leetcode.cn id=739 lang=golang
 *
 * [739] 每日温度
 */
package main

// @lc code=start
// 使用dp也不能生成n^2矩阵，这样复杂度和双重循环一样
// 使用单调栈去存储温度下标
// 每次新元素比较时，去比较单调栈中比自己温度低的，去更新
func dailyTemperatures(temperatures []int) []int {
	var idxStack []int
	ret := make([]int, len(temperatures))
	for idx, temp := range temperatures {
		for {
			if len(idxStack) == 0 ||
				temperatures[idxStack[len(idxStack)-1]] >= temp {
				break
			}
			gap := idx - idxStack[len(idxStack)-1]
			ret[idxStack[len(idxStack)-1]] = gap
			idxStack = idxStack[:len(idxStack)-1]
		}
		idxStack = append(idxStack, idx)
	}
	return ret
}

// 最简单的方法是双重遍历，复杂度是O(n^2)
func dailyTemperatures1(temperatures []int) []int {
	lens := len(temperatures)
	var ret []int
	for i := 0; i < lens; i++ {
		j := i
		for ; j < lens; j++ {
			if temperatures[j] > temperatures[i] {
				break
			}
		}
		if j == lens {
			j = i
		}
		ret = append(ret, j-i)
	}
	return ret
}

// @lc code=end
