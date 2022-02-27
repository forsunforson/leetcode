/*
 * @lc app=leetcode.cn id=1994 lang=golang
 *
 * [1994] 好子集的数目
 */
package main

// @lc code=start
var squareNums = []int{4, 8, 9, 12, 16, 18, 20, 24, 25, 27, 28}
var primeNums = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}

func tooslow(nums []int) int {
	// 有相同的因数，肯定不是好子集
	count := 0
	numsMap := make(map[int]byte)
	for _, num := range squareNums {
		numsMap[num] = 0
	}
	countOne := 0
	restNums := []int{}
	for _, num := range nums {
		if num == 1 {
			countOne++
			continue
		}
		if _, ok := numsMap[num]; ok {
			continue
		}
		restNums = append(restNums, num)
	}
	bitMap := make(map[int]uint16)
	for _, num := range restNums {
		bitMap[num] = sep(num)
	}
	var dfs func(uint16, int)
	dfs = func(u uint16, i int) {
		if i >= len(restNums) {
			return
		}
		if bitMap[restNums[i]]&u == 0 {
			count++
			dfs(bitMap[restNums[i]]|u, i+1)
		}
		dfs(u, i+1)
	}
	dfs(0, 0)
	for countOne > 0 {
		count *= 2
		count %= (1e9 + 7)
		countOne--
	}
	return count % (1e9 + 7)
}

func sep(num int) uint16 {
	var ret uint16 = 0
	for i := range primeNums {
		if num == 0 {
			break
		}
		if num%primeNums[i] == 0 {
			ret |= 1 << i
			num /= primeNums[i]
		}
	}
	return ret
}

func numberOfGoodSubsets(nums []int) (ans int) {
	const mod int = 1e9 + 7
	freq := [31]int{}
	for _, num := range nums {
		freq[num]++
	}

	f := make([]int, 1<<len(primeNums))
	f[0] = 1
	for i := 0; i < freq[1]; i++ {
		f[0] = f[0] * 2 % mod
	}
next:
	for i := 2; i < 31; i++ {
		if freq[i] == 0 {
			continue
		}

		// 检查 i 的每个质因数是否均不超过 1 个
		subset := 0
		for j, prime := range primeNums {
			if i%(prime*prime) == 0 {
				continue next
			}
			if i%prime == 0 {
				subset |= 1 << j
			}
		}

		// 动态规划
		for mask := 1 << len(primeNums); mask > 0; mask-- {
			if mask&subset == subset {
				f[mask] = (f[mask] + f[mask^subset]*freq[i]) % mod
			}
		}
	}

	for _, v := range f[1:] {
		ans = (ans + v) % mod
	}
	return
}

// @lc code=end
