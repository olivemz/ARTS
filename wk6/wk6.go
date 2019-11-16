package main

import (
	"fmt"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	intMax := len(nums) - 1
	intMin := 0
	intSum := nums[intMax] + nums[intMin] + nums[intMin+1]
	for intMin < intMax-1 {
		intFlexMax := intMax
		intFlex := intMin + 1
		for intFlex < intFlexMax {
			intTempSum := nums[intMin] + nums[intFlexMax] + nums[intFlex]
			if intTempSum == target {
				return intTempSum
			} else if intTempSum < target {
				for intFlex < intFlexMax && nums[intFlex] == nums[intFlex+1] {
					intFlex++
				}
				intFlex++
			} else {
				for intFlex < intFlexMax && nums[intFlexMax] == nums[intFlexMax-1] {
					intFlexMax--
				}
				intFlexMax--
			}
			if convertToCompare(intSum, intTempSum, target) {
					intSum = intTempSum
				}
		}
		intMin++
	}
	return intSum
}

func convertToCompare(x int, y int, target int) bool {
	xGap := x - target
	yGap := y - target
	if xGap < 0 {
		xGap = -xGap
	}
	if yGap < 0 {
		yGap = -yGap
	}
	if xGap > yGap {
		return true
	}
	return false
}
func main() {
	nums := []int{-1, 2, 1, -4}
	x := threeSumClosest(nums, 2)
	fmt.Println(x)
}
