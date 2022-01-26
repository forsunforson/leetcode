/*
 * @lc app=leetcode.cn id=10 lang=golang
 *
 * [10] 正则表达式匹配
 */
package main

// @lc code=start
func isMatch(s string, p string) bool {
	return isMatchCore(s, p, 0, 0)
}

func isMatchCore(s, p string, idxS, idxP int) bool {
	if idxS == len(s) && idxP == len(p) {
		// 都移动到结束字符，则匹配结束
		return true
	}
	if idxS > len(s) || idxP > len(p) {
		return false
	}
	if idxP < len(p) && p[idxP] == '.' {
		if idxP+1 < len(p) && p[idxP+1] == '*' {
			return isMatchCore(s, p, idxS, idxP+2) ||
				isMatchCore(s, p, idxS+1, idxP)
		}
		return isMatchCore(s, p, idxS+1, idxP+1)
	}
	if idxS < len(s) && idxP < len(p) && s[idxS] == p[idxP] {
		if idxP+1 < len(p) && p[idxP+1] == '*' {
			return isMatchCore(s, p, idxS, idxP+2) ||
				isMatchCore(s, p, idxS+1, idxP)
		} else {
			return isMatchCore(s, p, idxS+1, idxP+1)
		}
	}

	// 如果当前字符不等，看看后面是不是有*，有*继续判断，不然肯定不匹配
	if idxP+1 < len(p) && p[idxP+1] == '*' {
		return isMatchCore(s, p, idxS, idxP+2)
	}

	return false

}

// @lc code=end
