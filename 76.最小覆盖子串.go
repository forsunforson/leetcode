/*
 * @lc app=leetcode.cn id=76 lang=golang
 *
 * [76] 最小覆盖子串
 */
package main

// @lc code=start
func minWindow(s string, t string) string {
	targetDict := make(map[byte]int)
	for i := range t {
		targetDict[t[i]]++
	}
	dict := make(map[byte]int)
	minString := ""
	min := len(s) + 1
	p0, p1 := 0, 0
	for {
		//fmt.Println("curStr: ", s[p0:p1])
		if isCovered(targetDict, dict) {
			curStr := s[p0:p1]
			if min > len(curStr) {
				min = len(curStr)
				minString = curStr
			}
			dict[s[p0]]--
			p0++
		} else {
			if p1 == len(s) {
				break
			}
			dict[s[p1]]++
			p1++
		}

	}
	return minString
}

func isCovered(target, dict map[byte]int) bool {
	for k, v := range target {
		cur := dict[k]
		if v > cur {
			return false
		}
	}
	return true
}

// @lc code=end
