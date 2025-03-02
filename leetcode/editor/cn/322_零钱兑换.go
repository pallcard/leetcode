package main

import (
	"fmt"
	"math"
)

// leetcode submit region begin(Prohibit modification and deletion)

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	dpArr := make([]int, amount+1)

	for i := 0; i < len(dpArr); i++ {
		dpArr[i] = amount + 1
	}

	dpArr[0] = 0
	//for _, coin := range coins {
	//	dpArr[coin] = 1
	//}
	for i := 1; i < len(dpArr); i++ {

		for _, coin := range coins {
			if i-coin < 0 {
				continue
			}

			if dpArr[i] > dpArr[i-coin]+1 {
				dpArr[i] = dpArr[i-coin] + 1
			}
		}

	}

	if dpArr[amount] == amount+1 {
		return -1
	}
	return dpArr[amount]

	//memo := make([]int, amount+1)
	//for i := 0; i < len(memo); i++ {
	//	memo[i] = -2
	//}
	//return dp(coins, amount, memo)
}

func dp(coins []int, amount int, memo []int) int {
	if len(coins) == 0 {
		return -1
	}
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}

	if memo[amount] != -2 {
		return memo[amount]
	}

	min := math.MaxInt
	for i := 0; i < len(coins); i++ {
		rest := amount - coins[i]
		sub := dp(coins, rest, memo)
		if sub == -1 {
			continue
		}
		memo[rest] = sub
		if sub+1 < min {
			min = sub + 1
		}
	}
	if min == math.MaxInt {
		return -1
	}
	return min
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	change := coinChange([]int{2}, 1)
	fmt.Print(change)
}
