/*
 * @lc app=leetcode.cn id=468 lang=golang
 *
 * [468] 验证IP地址
 */
package main

import (
	"strconv"
	"strings"
)

// @lc code=start
func validIPAddress(queryIP string) string {
	if strings.Contains(queryIP, ".") {
		return validIPv4(queryIP)
	}
	if strings.Contains(queryIP, ":") {
		return validIPv6(queryIP)
	}
	return "Neither"
}

func validIPv4(ip string) string {
	dirs := strings.Split(ip, ".")
	if len(dirs) != 4 {
		return "Neither"
	}
	for _, dir := range dirs {
		n, err := strconv.Atoi(dir)
		if err != nil {
			return "Neither"
		}
		if n == 0 && len(dir) != 1 {
			return "Neither"
		}
		if n < 0 || n > 255 {
			return "Neither"
		}
		if n != 0 && dir[0] == '0' {
			return "Neither"
		}
	}
	return "IPv4"
}

func validIPv6(ip string) string {
	if ip[0] == ':' || ip[len(ip)-1] == ':' {
		return "Neither"
	}
	dirs := strings.Split(ip, ":")
	if len(dirs) > 8 {
		return "Neither"
	}
	for i, dir := range dirs {
		if len(dir) > 4 {
			return "Neither"
		}
		if len(dir) == 0 {
			if i > 1 && dirs[i-1] == "0" || i < len(dir)-1 && dirs[i+1] == "0" {
				return "Neither"
			}
		}
		for i := range dir {
			if dir[i] >= '0' && dir[i] <= '9' {
				continue
			}
			if dir[i] >= 'a' && dir[i] <= 'f' {
				continue
			}
			if dir[i] >= 'A' && dir[i] <= 'F' {
				continue
			}
			return "Neither"
		}
	}
	return "IPv6"
}

// @lc code=end
