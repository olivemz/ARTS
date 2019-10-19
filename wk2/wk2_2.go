package main
//https://leetcode.com/problems/coin-change
import (
	"fmt"
)

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount + 1)
	dp[0] = 0
	for x:=1; x< len(dp) ; x ++ {
		dp[x] = amount + 1
	}
	for i:=1; i<=amount; i++ {
		for j:= 0; j <= len(coins)-1 ; j++ {
			if coins[j] <= i {
				dp[i] = min( dp[i] , dp[i-coins[j]] + 1)
			}
		}
	}
	if dp[amount] > amount {
		return -1
	}

	return dp[amount]
}
func min(x int, y int) int {
	if x > y {
		return y
	}
	return x
}
func main(){
	x := []int{1}
	fmt.Print(coinChange(x, 6))
}