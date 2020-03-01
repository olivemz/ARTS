package main

import "fmt"

/**
Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Note that an empty string is also considered valid.
**/
func isValid(s string) bool {
	var left map[string]string
	left = make(map[string]string)
	left["("] = ")"
	left["{"] = "}"
	left["["] = "]"
	var stack []string
	stack = []string{}
	i := len(s)
	for j := 0; j < i; j++ {
		current := s[j : j+1]
		if val, ok := left[current]; ok {
			stack = append(stack, val)
		} else {
			if len(stack) == 0 {
				return false
			}
			pop := stack[len(stack)-1]
			stack = stack[0 : len(stack)-1]
			if pop != current {
				return false
			}
		}
	}
	if len(stack) > 0 {
		return false
	} 
	return true
}

func main() {
	fmt.Println(isValid("]"))
}
