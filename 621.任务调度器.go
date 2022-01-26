/*
 * @lc app=leetcode.cn id=621 lang=golang
 *
 * [621] 任务调度器
 */
package main

// @lc code=start
func leastInterval(tasks []byte, n int) int {
	taskMap := make(map[byte]int, 0)
	cd := make(map[byte]int, 0)
	total := 0
	time := 0
	// init taskMap
	for _, task := range tasks {
		count, ok := taskMap[task]
		if !ok {
			taskMap[task] = 1
		} else {
			taskMap[task] = count + 1
		}
		total++
	}
	// init cd
	for task := range taskMap {
		cd[task] = 0
	}
	//fmt.Printf("taskMap %v, cd %v\n", taskMap, cd)

	for total > 0 {
		curTask := selectTask(taskMap, cd)
		time++
		updateCD(cd, curTask, n)
		updateTaskMap(taskMap, curTask)
		if curTask != 0 {
			total--
		}
		//fmt.Printf("select %v, time %d\n", curTask, time)
	}
	return time
}

func updateTaskMap(taskMap map[byte]int, task byte) {
	if task == 0 {
		return
	}
	taskMap[task] = taskMap[task] - 1
}

func updateCD(cd map[byte]int, task byte, n int) {
	if task == 0 {
		for k, v := range cd {
			cd[k] = v - 1
		}
		return
	}

	for k, v := range cd {
		if k == task {
			cd[k] = n
		} else {
			if cd[k] == 0 {
				continue
			}
			cd[k] = v - 1
		}
	}

}

func selectTask(taskMap, cd map[byte]int) byte {
	// 选已cd中数量最多的任务
	var cdTasks []byte
	for k, v := range cd {
		if v == 0 {
			cdTasks = append(cdTasks, k)
		}
	}
	if len(cdTasks) == 0 {
		return 0
	}
	most := 0
	var retTask byte
	for _, task := range cdTasks {
		if taskMap[task] > most {
			retTask = task
			most = taskMap[task]
		}
	}
	return retTask
}

// @lc code=end
