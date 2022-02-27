/*
 * @lc app=leetcode.cn id=1360 lang=golang
 *
 * [1360] 日期之间隔几天
 */
package main

import (
	"strconv"
	"strings"
)

// @lc code=start
var monthDay = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

func daysBetweenDates(date1 string, date2 string) int {
	if date1 > date2 {
		return daysBetweenDates(date2, date1)
	}
	count := 0
	d1, d2 := parse(date1), parse(date2)
	for !equal(d1, d2) {
		d1 = addOne(d1)
		count++
	}
	return count
}

func equal(date1 []int, date2 []int) bool {
	for i := 0; i < 3; i++ {
		if date1[i] != date2[i] {
			return false
		}
	}
	return true
}

func parse(date string) []int {
	time := strings.Split(date, "-")
	year, _ := strconv.Atoi(time[0])
	month, _ := strconv.Atoi(time[1])
	day, _ := strconv.Atoi(time[2])
	return []int{year, month, day}
}

func addOne(date []int) []int {
	year, month, day := date[0], date[1], date[2]
	day++
	curMonth := monthDay[month-1]
	if month == 2 && (year%400 == 0 || year%100 != 0 && year%4 == 0) {
		curMonth = 29
	}
	if day > curMonth {
		month++
		day = 1
	}
	if month > 12 {
		year++
		month = 1
	}
	return []int{year, month, day}
}

// @lc code=end
