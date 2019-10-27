package main

import "fmt"

func maxArea(height []int) int {
	i := 0
	j := len(height) - 1
	maxSize := 0
	currSize := 0
	for i < j {
		currSize = returnArea(i , j , height)
		if currSize > maxSize {
			maxSize = currSize
		}
		leftHeight := height[i]
		rightHeight := height[j]
		if leftHeight < rightHeight {
			i = i + 1
		} else {
			j = j -1
		}
	}
	return maxSize
}

func returnArea(a int, b int , height []int) int {
	len := b - a
	if height[a] > height[b] {
		return height[b] * len
	}
	return height[a] * len
}

func main () {
	testArray := []int{1,2,3,4,5,25,24,3,4}
	result:= maxArea(testArray)
	fmt.Print(result)
}