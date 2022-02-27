/*
 * @lc app=leetcode.cn id=205 lang=golang
 *
 * [205] 同构字符串
 */

// @lc code=start
func isIsomorphic(s string, t string) bool {
	dict := map[byte]byte{}
	rDic := map[byte]byte{}
	for i := range s {
		v1, ok1 := dict[s[i]]
		v2, ok2 := rDic[t[i]]
		if ok1 != ok2 {
			return false
		}
		if !ok1 {
			dict[s[i]] = t[i]
			rDic[t[i]] = s[i]
		} else {
			if v1 != t[i] || s[i] != v2 {
				return false
			}
		}
	}
	return true
}

// @lc code=end

