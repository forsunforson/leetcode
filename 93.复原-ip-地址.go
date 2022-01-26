/*
 * @lc app=leetcode.cn id=93 lang=golang
 *
 * [93] 复原 IP 地址
 */
package main

import (
	"strconv"
	"strings"
)

// @lc code=start
/*
  思路：先全排列，再通过筛选函数，将有效的ip放到结果中
*/
func restoreIpAddresses(s string) []string {
	if len(s) > 12 {
		return []string{}
	}
	ret := make([]string, 0)
	getIpCore(s, "", 0, 3, 0, &ret)
	return ret
}

func getIpCore(s, curIp string, index, dot, seg int, ret *[]string) {
	if index == len(s) {
		if dot != 0 {
			return
		}
		if isValidIp(curIp) {
			*ret = append(*ret, curIp)
		}
		return
	}
	// 两种选择：1 加点 2 不加点
	curIp += string(s[index])
	if seg < 3 {
		getIpCore(s, curIp, index+1, dot, seg+1, ret)
	}
	if dot > 0 {
		newIp := curIp + "."
		getIpCore(s, newIp, index+1, dot-1, 0, ret)
	}
}
func isValidIp(s string) bool {
	ip := strings.Split(s, ".")
	if len(ip) != 4 {
		return false
	}
	for _, i := range ip {
		if len(i) < 1 {
			return false
		}
		num, _ := strconv.Atoi(i)
		if num < 0 || num > 255 {
			return false
		}
		if num != 0 {
			if i[0] == '0' {
				return false
			}
		} else {
			if len(i) > 1 {
				return false
			}
		}

	}
	return true
}

// @lc code=end
