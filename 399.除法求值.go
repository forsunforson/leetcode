/*
 * @lc app=leetcode.cn id=399 lang=golang
 *
 * [399] 除法求值
 */
package main

// @lc code=start
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	params := make(map[string]map[string]float64)
	// 把所有可能的组合算出来，然后和queries中进行判断
	for idx, equation := range equations {
		x, y := equation[0], equation[1]
		params[x] = map[string]float64{y: values[idx], x: 1}
		params[y] = map[string]float64{x: 1 / values[idx], y: 1}
	}
	ans := []float64{}
	for _, query := range queries {
		num := float64(1)
		if calDiv(params, query[0], query[1], &num) {
			ans = append(ans, num)
		} else {
			ans = append(ans, -1)
		}
	}
	return ans
}

func calDiv(params map[string]map[string]float64, source, target string, num *float64) bool {
	if _, ok := params[source]; !ok {
		return false
	}
	if _, ok := params[target]; !ok {
		return false
	}
	val, ok := params[source][target]
	if ok {
		*num = val * *num
		return true
	}
	for k, v := range params[source] {
		copyNum := *num * v
		if calDiv(params, k, target, &copyNum) {
			*num = copyNum
			return true
		}
	}
	return false
}

// @lc code=end
