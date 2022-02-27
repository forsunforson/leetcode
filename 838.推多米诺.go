/*
 * @lc app=leetcode.cn id=838 lang=golang
 *
 * [838] 推多米诺
 */
package main

// @lc code=start
func pushDominoes(dominoes string) string {
	// 对于L-R之间的牌，需要平分
	// 双指针
	size := len(dominoes)
	l, r := 0, size-1

	ans := []byte(dominoes)
	// 先处理左右两边空的
	for l < size && ans[l] == '.' {
		l++
	}
	if l == size {
		return string(ans)
	}
	if ans[l] == 'L' {
		for i := 0; i < l; i++ {
			ans[i] = 'L'
		}
	}
	for r >= 0 && ans[r] == '.' {
		r--
	}
	if ans[r] == 'R' {
		for i := r + 1; i < size; i++ {
			ans[i] = 'R'
		}
	}
	// 处理中间的
	for i := l; i < r; {
		for i < r && ans[i+1] != '.' {
			i++
		}
		j := i + 1
		for j <= r && ans[j] == '.' {
			j++
		}
		if j > r {
			break
		}
		if ans[i] == 'L' {
			if ans[j] == 'R' {
				i = j
				continue
			}
			if ans[j] == 'L' {
				for k := i + 1; k < j; k++ {
					ans[k] = 'L'
				}
				i = j
				continue
			}
		}
		if ans[i] == 'R' {
			if ans[j] == 'R' {
				for k := i + 1; k < j; k++ {
					ans[k] = 'R'
				}
				i = j
				continue
			}
			if ans[j] == 'L' {
				length := j - i - 1
				if length%2 == 1 {
					for k := i + 1; k < (j+i)/2; k++ {
						ans[k] = 'R'
					}
					for k := (j+i)/2 + 1; k < j; k++ {
						ans[k] = 'L'
					}
				} else {
					for k := i + 1; k <= (j+i)/2; k++ {
						ans[k] = 'R'
					}
					for k := (j+i)/2 + 1; k < j; k++ {
						ans[k] = 'L'
					}
				}
				i = j
			}
		}
	}
	return string(ans)
}

// @lc code=end
