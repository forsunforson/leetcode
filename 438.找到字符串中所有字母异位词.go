/*
 * @lc app=leetcode.cn id=438 lang=golang
 *
 * [438] 找到字符串中所有字母异位词
 */

// @lc code=start
func findAnagrams(s string, p string) []int {
	ans := []int{}
	// 滑动窗口
	sLen, pLen := len(s), len(p)
	if pLen > sLen {
		return ans
	}
	var sDic, pDic [26]int
	for i := 0; i < pLen; i++ {
		sDic[s[i]-'a']++
		pDic[p[i]-'a']++
	}
	if sDic == pDic {
		ans = append(ans, 0)
	}
	// 在s中滑动窗口找p的变形
	for i := 0; i < sLen-pLen; i++ {
		sDic[s[i]-'a']--
		sDic[s[i+pLen]-'a']++
		if sDic == pDic {
			ans = append(ans, i+1)
		}
	}
	return ans
}

// @lc code=end

