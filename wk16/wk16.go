package main

//https://leetcode.com/problems/generate-parentheses/solution/
import "fmt"

func generateParenthesis(n int) []string {
	ans := make([]string, 0)
	backTrack(&ans, "", 0, 0, n)
	return ans
}
func backTrack(ans *[]string, cur string, open int, close int, max int) {
	if len(cur) == 2*max {
		*ans = append(*ans, cur)
		return
	}
	if open < max {
		backTrack(ans, cur+"(", open+1, close, max)
	}
	if close < open {
		backTrack(ans, cur+")", open, close+1, max)
	}
}
func main() {
	x := generateParenthesis(2)
	fmt.Println(x)
}
