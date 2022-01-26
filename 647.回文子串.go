/*
 * @lc app=leetcode.cn id=647 lang=golang
 *
 * [647] 回文子串
 */
package main

// @lc code=start
func countSubstrings(s string) int {
	count := 0
	helper := "#"
	for _, c := range s {
		helper += string(c) + "#"
	}
	//fmt.Println(helper)
	for i := 0; i < len(helper); i++ {
		// 对于每个字符，以其为中心，向两边扩散
		count += countSubstringsHelper(helper, i)
	}
	return count
}

func countSubstringsHelper(s string, idx int) int {
	count := 0

	for i := 0; idx+i < len(s) && idx-i >= 0; i++ {
		if s[idx-i] == '#' || s[idx+i] == '#' {
			continue
		}
		if s[idx-i] == s[idx+i] {
			count++
		} else {
			break
		}
	}
	return count
}

// @lc code=end
