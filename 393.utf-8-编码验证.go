/*
 * @lc app=leetcode.cn id=393 lang=golang
 *
 * [393] UTF-8 编码验证
 */
package main

// @lc code=start
func validUtf8(data []int) bool {
	for i := 0; i < len(data); {
		head := data[i]

		mask := 0x00000080
		if head&mask == 0 {
			// 只有一字节
			i++
		} else {
			digit := 0
			for head&mask != 0 {
				digit++
				mask >>= 1
			}
			if digit > 4 || digit == 1 {
				return false
			}
			for j := 1; j < digit; j++ {
				if i+j >= len(data) {
					return false
				}
				// 检查是不是10开头
				body := data[i+j]
				if body&128 != 128 {
					return false
				}
			}
			i += digit
		}

	}
	return true
}

// @lc code=end
