package main

// leetcode submit region begin(Prohibit modification and deletion)
func maxProfit2(prices []int) int {
	// todo dp[i][j][k] ，(k=0不持有股票，1持有股票)第i天交易上限为k最大利润, 以buy算交易上限
	// todo dp[i][j][0] =  max ( dp[i-1][j][0] 「不操作」, dp[i-1][j][1] + p[i] 「今天卖出股票」)
	// todo dp[i][j][1] =  max ( dp[i-1][j][1] 「不操作」, dp[i-1][j-1][0] - p[i] 「今天买入股票」)
	// todo 结果返回 dp[len(p)-1][K][0]
	// todo 初态 dp[0][j][0] = 0   , dp[0][j][1] = -p[0]
	// todo     dp[i][0][k] = 0

	// K == 1
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
		dp[i][1] = max(dp[i-1][1], -prices[i]) //todo 仅交易一次，之前未买过股票，不可能有收益即dp[i-1][0]=0
	}
	return dp[len(prices)-1][0]
}

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	// todo 压缩空间
	dp_i_0, dp_i_1 := 0, -prices[0]
	for i := 0; i < len(prices); i++ {
		dp_i_0 = max(dp_i_0, dp_i_1+prices[i])
		dp_i_1 = max(dp_i_1, -prices[i])
	}
	return dp_i_0

}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
