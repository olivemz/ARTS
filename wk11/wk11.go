package main

import "fmt"

var m = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

func backTrackLetter(combinations string, nextString string, result *[]string) {
	if len(nextString) == 0 {
		*result = append(*result, combinations)
	} else {
		digit := nextString[0:1]
		possibleString := m[digit]
		for _, val := range possibleString {
			backTrackLetter(combinations+string(val), nextString[1:], result)
		}
	}
}

func letterCombinations(digits string) []string {
	var result []string
	if len(digits) != 0 {
		backTrackLetter("", digits, &result)
	}
	return result
}

func main() {
	digits := "32"
	cominations := letterCombinations(digits)
	fmt.Printf("%v", cominations)
}
