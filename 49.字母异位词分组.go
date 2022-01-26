/*
 * @lc app=leetcode.cn id=49 lang=golang
 *
 * [49] 字母异位词分组
 */
package main

import "fmt"

// @lc code=start
func groupAnagrams(strs []string) [][]string {
	// 关键时得找到如何判断两个词是异位词
	//
	results := make([][]string, 0)
	strMap := make(map[string][]string)
	for _, str := range strs {
		key := getKey(str)
		result := strMap[key]
		result = append(result, str)
		strMap[key] = result
	}
	for _, v := range strMap {
		results = append(results, v)
	}
	return results
}

func getKey(s string) string {
	count := make([]int, 26)
	for i := 0; i < len(s); i++ {
		count[s[i]-'a']++
	}
	key := ""
	for r, c := range count {
		if c != 0 {
			key += string(rune(r) + 'a')
			key += fmt.Sprint(c)
		}
	}
	return key
}

// @lc code=end
