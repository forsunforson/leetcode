/*
 * @lc app=leetcode.cn id=1189 lang=golang
 *
 * [1189] “气球” 的最大数量
 */
package main

// @lc code=start
func maxNumberOfBalloons(text string) int {
	wordCount := [26]int{}
	for i := 0; i < len(text); i++ {
		wordCount[text[i]-'a']++
	}
	balloon := map[int]int{
		'b' - 'a': 1,
		0:         1,
		'l' - 'a': 2,
		'o' - 'a': 2,
		'n' - 'a': 1,
	}
	count := 0
	for {
		for k, v := range balloon {
			if wordCount[k] < v {
				return count
			}
			wordCount[k] -= v
		}
		count++
	}
}

// @lc code=end
