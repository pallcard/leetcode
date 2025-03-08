package main
//leetcode submit region begin(Prohibit modification and deletion)
func maxProfit(k int, prices []int) int {

	if len(prices) <= 1 {
		return 0
	}


	dp := make([][][2]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([][2]int, k)
	}

	for i := 0; i < len(prices); i++ {
		for j := k-1; j >= 0; j-- {
			if i == 0 {
				dp[i][j][0] = 0
				dp[i][j][1] = -prices[i]
				continue
			}

			dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j][1] + prices[i])
			if j == 0 {
				dp[i][j][1] = max(dp[i-1][j][1], -prices[i])
			} else {
				dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j-1][0]-prices[i])
			}


		}
	}


	return dp[len(prices)-1][k-1][0]
}
//leetcode submit region end(Prohibit modification and deletion)


func main() {
	
}

