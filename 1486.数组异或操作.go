/*
 * @lc app=leetcode.cn id=1486 lang=golang
 *
 * [1486] 数组异或操作
 */
package main

// @lc code=start
/*
  对于n为偶数的输入，start的偶数次XOR为0；
  奇数，XOR为本身。
  1 0;2 2;3 6;4 0;5 8;6 2;7
*/
func xorOperation(n int, start int) int {
	ret := start
	for i := 1; i < n; i++ {
		ret = ret ^ (start + 2*i)
	}
	return ret
}

// @lc code=end
