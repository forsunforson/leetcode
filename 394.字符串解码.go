/*
 * @lc app=leetcode.cn id=394 lang=golang
 *
 * [394] 字符串解码
 */
package main

// @lc code=start
func decodeString(s string) string {
	return decodeCore(s)
}

// 输入k[string]的形式，递归调用
func decodeCore(s string) string {
	ret := ""

	for idx := 0; idx < len(s); {
		num := 0
		if isDigit(rune(s[idx])) {
			num = int(s[idx] - '0')
			for {
				idx++
				if !isDigit(rune(s[idx])) {
					break
				}
				num = num*10 + int(s[idx]-'0')
			}
			nextIdx := findSubStr(s, idx)
			repeatStr := decodeCore(s[idx+1 : nextIdx])
			for repeat := 0; repeat < num; repeat++ {
				ret += repeatStr
			}
			idx = nextIdx + 1
		} else {
			ret += string(rune(s[idx]))
			idx++
		}
	}
	return ret
}

func findSubStr(s string, idx int) int {
	count := 1
	i := idx + 1
	for ; i < len(s); i++ {
		if s[i] == byte('[') {
			count++
		}
		if s[i] == byte(']') {
			count--
		}
		if count == 0 {
			break
		}
	}
	return i
}

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

// @lc code=end
