/*
 * @lc app=leetcode.cn id=67 lang=golang
 *
 * [67] 二进制求和
 */

// @lc code=start
func addBinary(a string, b string) string {
	if len(b) > len(a) {
		return addBinary(b, a)
	}
	ans := make([]byte, len(a)+1)
	c := byte(0)
	for i := 1; i <= len(b); i++ {
		sum := a[len(a)-i] - '0' + b[len(b)-i] - '0' + c
		if sum > 1 {
			sum -= 2
			c = 1
		} else {
			c = 0
		}
		ans[len(ans)-i] = sum + '0'
	}
	for i := len(b) + 1; i <= len(a); i++ {
		sum := a[len(a)-i] - '0' + c
		if sum > 1 {
			sum -= 2
			c = 1
		} else {
			c = 0
		}
		ans[len(ans)-i] = sum + '0'
	}
	if c == 1 {
		ans[0] = 1 + '0'
		return string(ans)
	}
	return string(ans[1:])
}

// @lc code=end

