/*
 * @lc app=leetcode.cn id=207 lang=golang
 *
 * [207] 课程表
 */
package main

import "fmt"

// @lc code=start
func canFinish(numCourses int, prerequisites [][]int) bool {
	// 从每个点作为起点，深度优先遍历
	coursesMap := make(map[int]int)
	prerequisitesMap := parsePrerequisites(prerequisites)
	for i := 0; i < numCourses; i++ {
		coursesMap[i] = 0
	}
	for i := 0; i < numCourses; i++ {
		result := dfsHelper(i, &coursesMap, &prerequisitesMap)
		if !result {
			return false
		}
	}
	fmt.Printf("map %v\n", coursesMap)
	for _, v := range coursesMap {
		if v != 1 {
			return false
		}
	}
	return true
}

func dfsHelper(curCourse int, coursesMap *map[int]int, prerequisites *map[int][]int) bool {
	fmt.Printf("cur %d ", curCourse)
	if (*coursesMap)[curCourse] == 1 {
		return true
	}
	for {
		nextCourse := getNextCourse(curCourse, prerequisites)
		if nextCourse == -1 {
			break
		}
		fmt.Printf("next %d \n", nextCourse)
		(*coursesMap)[curCourse] = -1
		if (*coursesMap)[nextCourse] == -1 {
			return false
		}
		ret := dfsHelper(nextCourse, coursesMap, prerequisites)
		if ret {
			(*coursesMap)[curCourse] = 1
			continue
		} else {
			return false
		}
	}
	(*coursesMap)[curCourse] = 1

	return true
}

func parsePrerequisites(prerequisites [][]int) map[int][]int {
	ret := make(map[int][]int)
	for _, prerequisite := range prerequisites {
		courses, ok := ret[prerequisite[0]]
		if !ok {
			ret[prerequisite[0]] = []int{prerequisite[1]}
		} else {
			ret[prerequisite[0]] = append(courses, prerequisite[1])
		}
	}
	return ret
}

func getNextCourse(curCourse int, prerequisites *map[int][]int) int {
	nextCourse := -1
	nextCourses, ok := (*prerequisites)[curCourse]
	if ok {
		if len(nextCourses) > 0 {
			nextCourse = nextCourses[0]
			nextCourses = nextCourses[1:]
			(*prerequisites)[curCourse] = nextCourses
		}
	}
	return nextCourse
}

// @lc code=end
