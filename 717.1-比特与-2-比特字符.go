/*
 * @lc app=leetcode.cn id=717 lang=golang
 *
 * [717] 1比特与2比特字符
 */
package main

// @lc code=start
func isOneBitCharacter(bits []int) bool {
	// 遇到1走两步，遇到0走一步
	// 听到最后一个位置说明是一字符结尾
	ptr := 0
	for ptr < len(bits)-1 {
		if bits[ptr] == 0 {
			ptr++
		} else {
			ptr += 2
		}
	}
	return ptr == len(bits)-1
}

func backword(bits []int) bool {
	// 倒序
	idx := len(bits) - 2
	for idx >= 0 {
		if bits[idx] != 0 {
			idx--
		}
	}
	countOne := len(bits) - 1 - idx - 1
	return countOne%2 == 0
}

// @lc code=end
