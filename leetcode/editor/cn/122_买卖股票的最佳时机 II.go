package main

// leetcode submit region begin(Prohibit modification and deletion)
func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	dp := make([][2]int, len(prices))

	for i := 0; i < len(prices); i++ {
		if i == 0 {
			dp[i][0] = 0
			dp[i][1] = -prices[i]
			continue
		}

		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i]) //todo 可无限制交易
	}

	return dp[len(prices)-1][0]

}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
