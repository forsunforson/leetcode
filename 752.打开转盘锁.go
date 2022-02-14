/*
 * @lc app=leetcode.cn id=752 lang=golang
 *
 * [752] 打开转盘锁
 */
package main

// @lc code=start

func openLock(deadends []string, target string) int {
	if target == "0000" {
		return 0
	}
	// 回溯法，一个个试，但是没法找出最优解
	// BFS通常用来找最优解
	visited := make(map[string]bool)
	for _, deadend := range deadends {
		visited[deadend] = true
	}
	if visited["0000"] {
		return -1
	}
	steps := make(map[string]int)
	queue := []string{}
	queue = append(queue, "0000")
	steps["0000"] = 0
	visited["0000"] = true
	for len(queue) != 0 {
		cur := queue[0]
		queue = queue[1:]
		for _, str := range generate(cur) {
			if !visited[str] {
				if str == target {
					return steps[cur] + 1
				}
				visited[str] = true
				queue = append(queue, str)
				steps[str] = steps[cur] + 1
			}
		}
	}
	return -1
}

func generate(cur string) []string {
	ans := []string{}
	for i := 0; i < 4; i++ {
		upstr := cur[:i]
		downstr := upstr
		if cur[i] == '0' {
			downstr += "9"
			upstr += string(cur[i] + 1)
		} else if cur[i] == '9' {
			downstr += string(cur[i] - 1)
			upstr += "0"
		} else {
			downstr += string(cur[i] - 1)
			upstr += string(cur[i] + 1)
		}
		for j := i + 1; j < 4; j++ {
			upstr += string(cur[j])
			downstr += string(cur[j])
		}
		ans = append(ans, upstr, downstr)
	}
	return ans
}

func rowDown(cur string, idx int) string {
	ans := cur[:idx]
	if cur[idx] == '0' {
		ans += "9"
	} else {
		ans += string(cur[idx] - 1)
	}
	for i := idx + 1; i < len(cur); i++ {
		ans += string(cur[i])
	}
	return ans
}

func rowUp(cur string, idx int) string {
	ans := cur[:idx]
	if cur[idx] == '9' {
		ans += "0"
	} else {
		ans += string(1 + cur[idx])
	}
	for i := idx + 1; i < len(cur); i++ {
		ans += string(cur[i])
	}
	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// @lc code=end
