package main

// Design a hit counter which counts the number of hits received in the past 5 minutes.
// 假定Gethit后，不会再hit 3分钟之前的时间戳了
type HitCounter struct {
	times []int
	hits  []int
}

func NewHitCounter() HitCounter {
	return HitCounter{
		times: make([]int, 300),
		hits:  make([]int, 300),
	}
}

func (h *HitCounter) Hit(t int) {
	timeIdx := t % 300
	if h.times[timeIdx] == t {
		h.hits[timeIdx]++
	} else {
		h.hits[timeIdx] = 1
		h.times[timeIdx] = t
	}
}

func (h *HitCounter) GetHit(t int) int {
	sum := 0
	for idx, time := range h.times {
		if t-time < 300 {
			sum += h.hits[idx]
		}
	}
	return sum
}
