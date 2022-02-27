/*
 * @lc app=leetcode.cn id=126 lang=golang
 *
 * [126] 单词接龙 II
 */
package main

import "math"

// @lc code=start
func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	find := false
	for _, word := range wordList {
		if word == endWord {
			find = true
		}
	}
	if !find {
		return [][]string{}
	}
	type pair struct {
		word    string
		step    int
		words   []string
		visited map[string]bool
	}
	ans := [][]string{}
	count := math.MaxInt32
	// 寻找最短路径，BFS，需要记录已经访问过的
	q := []pair{{beginWord, 1, []string{beginWord}, map[string]bool{}}}
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		if cur.word == endWord {
			if count == math.MaxInt32 || count == cur.step {
				count = cur.step
				ans = append(ans, cur.words)
			}
			continue
		}
		cur.visited[cur.word] = true
		for _, word := range wordList {
			if !cur.visited[word] && diff(cur.word, word) == 1 {
				newWords := append([]string{}, cur.words...)
				newWords = append(newWords, word)
				p := pair{
					word:    word,
					step:    cur.step + 1,
					words:   newWords,
					visited: cur.visited,
				}
				q = append(q, p)
			}
		}
	}
	return ans
}

func diff(s1, s2 string) int {
	n := 0
	for i := range s1 {
		if s1[i] != s2[i] {
			n++
		}
	}
	return n
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// @lc code=end
