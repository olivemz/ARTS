package main

import (
	"fmt"
)

func threeSum(nums []int) [][]int {
	// sort.Ints(nums)
	// var m = make(map[int]int)
	// var result = make([][]int, 0)
	// for _, val := range nums {
	// 	val, exist := m[val]
	// 	if exist {
	// 		m[val] = m[val] + 1
	// 	} else {
	// 		m[val] = 1
	// 	}
	// }
	// curMaxIndex := len(nums) - 1
	// remainVal := 0
	// for i := 0; i < curMaxIndex-1; i++ {
	// 	if i < curMaxIndex-1 && nums[i+1] == nums[i] {
	// 		continue
	// 	}

	// 	remainVal = nums[i]
	// 	for j := curMaxIndex; j > i; j-- {
	// 		if j > i+1 && nums[j-1] == nums[j] {
	// 			continue
	// 		}
	// 		testVal := remainVal
	// 		testVal += nums[j]
	// 		lastVal := 0 - testVal
	// 		if lastVal < nums[i] {
	// 			continue
	// 		}
	// 		repeatTime := 1
	// 		if lastVal == nums[j] {
	// 			repeatTime++
	// 		}
	// 		if lastVal == nums[i] {
	// 			repeatTime++
	// 		}
	// 		val, exist := m[lastVal]
	// 		if exist && val >= repeatTime {
	// 			result = append(result, []int{nums[i], nums[j], lastVal})
	// 		}
	// 	}
	// }

	// return result
}

func main() {
	x := threeSum([]int{-1, 0, 1, 2, -1, -4})
	fmt.Println(x)

}
