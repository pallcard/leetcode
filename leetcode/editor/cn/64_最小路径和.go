package main

// leetcode submit region begin(Prohibit modification and deletion)
func minPathSum(grid [][]int) int {
	//todo dp[i][j] 表示冲 (0,0) -> (i,j) 最小路径和， 求dp[m-1][n-1]

	dp := make([][]int, len(grid), len(grid))
	for i := 0; i < len(grid); i++ {
		dp[i] = make([]int, len(grid[0]), len(grid[0]))
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if i == 0 && j == 0 {
				dp[i][j] = grid[i][j]
			} else if i-1<0 {
				dp[i][j] = dp[i][j-1] + grid[i][j]
			}else if j-1<0 {
				dp[i][j] = dp[i-1][j] + grid[i][j]
			} else if dp[i-1][j] < dp[i][j-1] {
				dp[i][j] = dp[i-1][j] + grid[i][j]
			} else {
				dp[i][j] = dp[i][j-1] + grid[i][j]
			}

		}
	}
	return dp[len(grid)-1][len(grid[0])-1]
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
