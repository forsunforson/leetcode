/*
 * @lc app=leetcode.cn id=57 lang=golang
 *
 * [57] 插入区间
 */
package main

// @lc code=start
func insert(intervals [][]int, newInterval []int) [][]int {
	var ret [][]int
	if len(intervals) == 0 {
		ret = append(ret, newInterval)
		return ret
	}
	used := false
	if intervals[0][0] > newInterval[1] {
		ret = append(ret, newInterval)
		ret = append(ret, intervals...)
		return ret
	}

	for i := 0; i < len(intervals); {
		if intervals[i][0] > newInterval[1] && !used {
			ret = append(ret, newInterval)
			used = true
		} else if intervals[i][1] < newInterval[0] || intervals[i][0] > newInterval[1] {
			ret = append(ret, intervals[i])
			i++
		} else {
			used = true
			j := i + 1
			var curStart, curEnd int = intervals[i][0], intervals[i][1]
			if newInterval[0] < curStart {
				curStart = newInterval[0]
			}
			if newInterval[1] > curEnd {
				curEnd = newInterval[1]
			}
			for j < len(intervals) {
				if newInterval[1] >= intervals[j][0] {
					if intervals[j][1] > newInterval[1] {
						curEnd = intervals[j][1]
					} else {
						curEnd = newInterval[1]
					}
					j++
				} else {
					break
				}
			}
			ret = append(ret, []int{curStart, curEnd})
			i = j
		}
	}
	if !used {
		ret = append(ret, newInterval)
	}

	return ret
}

// @lc code=end
