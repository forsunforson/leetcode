/*
 * @lc app=leetcode.cn id=301 lang=golang
 *
 * [301] 删除无效的括号
 */
package main

// @lc code=start
func removeInvalidParentheses(s string) []string {
	// 初看的思路是dp，但是如何确认是最少的括号？
	// 遍历，然后把s作为key，删除括号数作为value？
	// BFS
	curStr := map[string]bool{s: true}
	ans := []string{}
	for {
		for k := range curStr {
			if isValid(k) {
				ans = append(ans, k)
			}
		}
		if len(ans) > 0 {
			return ans
		}
		newStr := make(map[string]bool)
		for k := range curStr {
			for i := 0; i < len(k); i++ {
				if i != 0 && k[i] == k[i-1] {
					continue
				}
				if k[i] != '(' && k[i] != ')' {
					continue
				}
				str := k[:i]
				str += k[i+1:]
				newStr[str] = true
			}
		}
		curStr = newStr
	}
}

func isValid(s string) bool {
	l := 0
	for _, c := range s {
		if c == '(' {
			l++
		} else if c == ')' {
			l--
			if l < 0 {
				return false
			}
		}
	}
	return l == 0
}

// @lc code=end
