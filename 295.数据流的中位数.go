/*
 * @lc app=leetcode.cn id=295 lang=golang
 *
 * [295] 数据流的中位数
 */
package main

import (
	"container/heap"
	"sort"
)

// @lc code=start
type MedianFinder struct {
	queMin, queMax hp
}

func Constructor() MedianFinder {
	return MedianFinder{}
}

func (mf *MedianFinder) AddNum(num int) {
	minQ, maxQ := &mf.queMin, &mf.queMax
	if minQ.Len() == 0 || num <= -minQ.IntSlice[0] {
		heap.Push(minQ, -num)
		if maxQ.Len()+1 < minQ.Len() {
			heap.Push(maxQ, -heap.Pop(minQ).(int))
		}
	} else {
		heap.Push(maxQ, num)
		if maxQ.Len() > minQ.Len() {
			heap.Push(minQ, -heap.Pop(maxQ).(int))
		}
	}
}

func (mf *MedianFinder) FindMedian() float64 {
	minQ, maxQ := mf.queMin, mf.queMax
	if minQ.Len() > maxQ.Len() {
		return float64(-minQ.IntSlice[0])
	}
	return float64(maxQ.IntSlice[0]-minQ.IntSlice[0]) / 2
}

type hp struct{ sort.IntSlice }

func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
// @lc code=end

// type MedianFinder struct {
// 	nums   []int
// 	count  int
// 	offset int
// }

// func Constructor() MedianFinder {
// 	return MedianFinder{nums: make([]int, 201), count: 0, offset: 100}
// }

// func (this *MedianFinder) AddNum(num int) {
// 	this.nums[num+this.offset]++
// 	this.count++
// }

// func (this *MedianFinder) FindMedian() float64 {
// 	if this.count%2 == 1 {
// 		target := this.count/2 + 1
// 		for n, c := range this.nums {
// 			if c == 0 {
// 				continue
// 			}
// 			if target <= c {
// 				return float64(n - this.offset)
// 			}
// 			target -= c
// 		}
// 	} else {
// 		t1 := this.count / 2
// 		t2 := t1 + 1
// 		n1, n2 := 0, 0
// 		for n, c := range this.nums {
// 			if c == 0 {
// 				continue
// 			}
// 			if t1 > 0 && t1 <= c {
// 				n1 = n
// 			}
// 			t1 -= c

// 			if t2 > 0 && t2 <= c {
// 				n2 = n
// 			}
// 			t2 -= c

// 			if t1 < 0 && t2 < 0 {
// 				break
// 			}
// 		}
// 		return float64(n1+n2-2*this.offset) / 2
// 	}
// 	return 0
// }
