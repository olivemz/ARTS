package main

import (
	"fmt"
	"sort"
)

// Given an array nums of n integers and an integer target, are there elements a, b, c, and d in nums such that a + b + c + d = target? Find all unique quadruplets in the array which gives the sum of target.

// Note:

// The solution set must not contain duplicate quadruplets.

// Example:

// Given array nums = [1, 0, -1, 0, -2, 2], and target = 0.

// A solution set is:
// [
//   [-1,  0, 0, 1],
//   [-2, -1, 1, 2],
//   [-2,  0, 0, 2]
// ]

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	m := make(map[int]int)
	result := make([][]int, 0)
	for _, v := range nums {
		m[v]++
	}
	maxIndex := len(nums) - 1
	lastFound := false
	lastFound2 := false
	lastFound3 := false
	for i := 0; i < maxIndex-2; i++ {
		threesumTargets := target - nums[i]
		if lastFound == true && i >= 1 && nums[i] == nums[i-1] {
			continue
		}
		lastFound = false

		for j := i + 1; j <= maxIndex-1; j++ {
			if lastFound2 == true && j >= 1 && nums[j] == nums[j-1] {
				continue
			}
			lastFound2 = false
			leftTarget := threesumTargets - nums[j]
			for x := maxIndex; x > j; x-- {
				if lastFound3 == true && x < maxIndex-1 && nums[x+1] == nums[x] {
					continue
				}
				lastFound3 = false
				lookUp := leftTarget - nums[x]
				if lookUp < nums[j] || lookUp > nums[x] {
					continue
				}
				val, exist := m[lookUp]
				//fmt.Println(nums[i], nums[j], nums[x], val)
				repeat := 1
				if exist {
					if lookUp == nums[i] {
						repeat++
					}
					if lookUp == nums[j] {
						repeat++
					}
					if lookUp == nums[x] {
						repeat++
					}
					if repeat <= val {
						result = append(result, []int{nums[i], nums[j], nums[x], lookUp})
						lastFound = true
						lastFound2 = true
						lastFound3 = true
					}
				}
			}
		}
	}
	if m[target] >= 4 {
		result = append(result, []int{target, target, target, target})
	}

	return result
}

func main() {
	x := fourSum([]int{0, 1, 5, 0, 1, 5, 5, -4}, 11)
	fmt.Println(x)
}

// TBC...
