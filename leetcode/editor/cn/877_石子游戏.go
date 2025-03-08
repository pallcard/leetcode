package main

// leetcode submit region begin(Prohibit modification and deletion)
func stoneGame(piles []int) bool {
	//todo 同486

	if len(piles) == 1 {
		return true
	}

	dp := make([][][]int, len(piles))
	for i := 0; i < len(piles); i++ {
		dp[i] = make([][]int, len(piles))
		for j := 0; j < len(piles); j++ {
			dp[i][j] = make([]int, 2)
		}
	}

	for i := 0; i < len(piles); i++ {
		dp[i][i][0] = piles[i]
		dp[i][i][0] = 0
	}

	for i := len(piles) - 2; i >= 0; i-- {
		for j := i + 1; j < len(piles); j++ {
			left := dp[i+1][j][1]+piles[i]
			right := dp[i][j-1][1]+piles[j]
			if left > right {
				dp[i][j][0] = left
				dp[i][j][1] = dp[i+1][j][0]
			} else {
				dp[i][j][0] = right
				dp[i][j][1] = dp[i][j-1][0]
			}
		}

	}


	return dp[0][len(piles)-1][0] >= dp[0][len(piles)-1][1]
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
