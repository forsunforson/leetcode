/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合
 */
package main

// @lc code=start
func letterCombinations(digits string) []string {
	digMap := map[rune][]rune{
		'0': {},
		'1': {},
		'2': {'a', 'b', 'c'},
		'3': {'d', 'e', 'f'},
		'4': {'g', 'h', 'i'},
		'5': {'j', 'k', 'l'},
		'6': {'m', 'n', 'o'},
		'7': {'p', 'q', 'r', 's'},
		'8': {'t', 'u', 'v'},
		'9': {'w', 'x', 'y', 'z'},
	}
	result := make([]string, 0)
	if len(digits) < 1 {
		return result
	}
	combHelper(digits, 0, digMap, "", &result)
	return result
}

func combHelper(digits string, cur int, digMap map[rune][]rune, str string, result *[]string) {
	if cur == len(digits) {
		(*result) = append((*result), str)
		return
	}
	curDig := digits[cur]
	runes := digMap[rune(curDig)]
	for _, r := range runes {
		newStr := str + string(r)
		combHelper(digits, cur+1, digMap, newStr, result)
	}
}

// @lc code=end
