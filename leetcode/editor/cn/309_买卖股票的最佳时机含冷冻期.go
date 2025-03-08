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
		if i == 1 {
			dp[i][1] = max(dp[i-1][1], -prices[i])
		} else {
			dp[i][1] = max(dp[i-1][1], dp[i-2][0]-prices[i]) //todo 买前两天的股票
			// todo 假设从i来买转移 dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])  (1)
			// todo i-1买来源 dp[i-1][0] = max(dp[i-2][0], dp[i-2][1]+prices[i])
			// todo 要使得（1）成立， i-1不能卖出，则 dp[i-1][0] = dp[i-2][0]
		}

	}

	return dp[len(prices)-1][0]

}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
