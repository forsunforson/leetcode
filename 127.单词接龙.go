/*
 * @lc app=leetcode.cn id=127 lang=golang
 *
 * [127] 单词接龙
 */
package main

// @lc code=start
func ladderLength(beginWord string, endWord string, wordList []string) int {
	find := false
	for _, word := range wordList {
		if word == endWord {
			find = true
		}
	}
	if !find {
		return 0
	}
	type pair struct {
		word string
		step int
	}
	// 寻找最短路径，BFS，需要记录已经访问过的
	q := []pair{{beginWord, 1}}
	visited := make(map[string]bool)
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		if visited[cur.word] {
			continue
		}
		if cur.word == endWord {
			return cur.step
		}
		visited[cur.word] = true
		for _, word := range wordList {
			if diff(cur.word, word) == 1 {
				q = append(q, pair{word, cur.step + 1})
			}
		}
	}
	return 0
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
