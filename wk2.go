package main

import (
	"bytes"
	"fmt"
)

var blnDown bool

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	x := make([][]string, numRows)
	charPos := 0
	whichRow := 0
	blnDown = true
	var result bytes.Buffer
	for charPos < len(s) {
		x[whichRow] = append(x[whichRow], s[charPos:charPos+1])
		if blnDown {
			whichRow++
		} else {
			whichRow--
		}

		if whichRow-1 < 0 {
			blnDown = true
		} else if whichRow+1 > numRows-1 {
			blnDown = false
		}
		charPos++
	}
	row := 0
	column := 0
	for row < numRows {
		column = 0
		for column < len(x[row]) {
			result.WriteString(x[row][column])
			column++
		}
		row++
	}
	return result.String()
}

func main() {
	result := convert("AB", 1)
	fmt.Print(result)
}
