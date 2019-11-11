package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var m = make(map[int]int)
	var result = make([][]int, 0)
	for _, val := range nums {
		_, exist := m[val]
		if exist {
			m[val] = m[val] + 1
		} else {
			m[val] = 1
		}
	}
	curMaxIndex := len(nums) - 1
	remainVal := 0
	j := curMaxIndex
	for i := 0; i < curMaxIndex-1 && i < j; i++ {
		if i < curMaxIndex-1 && nums[i+1] == nums[i] {
			continue
		}

		remainVal = nums[i]
		for ; j > i; j-- {
			if j > i+1 && nums[j-1] == nums[j] {
				continue
			}
			lastVal := 0 - nums[j] - remainVal
			if lastVal < nums[i] || lastVal > nums[j] {
				continue
			}
			repeatTime := 1
			if lastVal == nums[j] {
				repeatTime++
			}
			if lastVal == nums[i] {
				repeatTime++
			}
			val, exist := m[lastVal]
			if exist && val >= repeatTime {
				result = append(result, []int{nums[i], nums[j], lastVal})
			}
		}
	}
	if m[0] >= 3 {
		result = append(result, []int{0, 0, 0})
	}
	return result
}

func main() {
	x := threeSum([]int{0, 0, 0})
	fmt.Println(x)

}
