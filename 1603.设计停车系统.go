/*
 * @lc app=leetcode.cn id=1603 lang=golang
 *
 * [1603] 设计停车系统
 */
package main

// @lc code=start
type ParkingSystem struct {
	big, medium, small                   int
	parkedBig, parkedMedium, parkedSmall int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	p := ParkingSystem{
		big:    big,
		medium: medium,
		small:  small,
	}
	return p
}

func (this *ParkingSystem) AddCar(carType int) bool {
	switch carType {
	case 1:
		if this.parkedBig < this.big {
			this.parkedBig++
			return true
		}
		return false
	case 2:
		if this.parkedMedium < this.medium {
			this.parkedMedium++
			return true
		}
		return false
	case 3:
		if this.parkedSmall < this.small {
			this.parkedSmall++
			return true
		}
		return false
	default:
		return false
	}
}

/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */
// @lc code=end
