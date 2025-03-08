package main

import "fmt"

// leetcode submit region begin(Prohibit modification and deletion)
func calculateMinimumHP(dungeon [][]int) int {
	// todo dp[i][j]:  (i,j)->(m-1,n-1)的最少血量

	m := len(dungeon)
	n := len(dungeon[0])

	dp := make([][]int, m, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == m-1 && j == n-1 { // m-1, n-1
				if dungeon[m-1][n-1] <= 0 {
					dp[m-1][n-1] = -dungeon[m-1][n-1] + 1
				} else {
					dp[m-1][n-1] = 1
				}
			} else if i+1 >= m { //看右边数
				if dp[i][j+1]-dungeon[i][j] > 0 {
					dp[i][j] = dp[i][j+1] - dungeon[i][j]
				} else {
					dp[i][j] = 1
				}
			} else if j+1 >= n { //看下面数
				if dp[i+1][j]-dungeon[i][j] > 0 {
					dp[i][j] = dp[i+1][j] - dungeon[i][j]
				} else {
					dp[i][j] = 1
				}
			} else { //看右边、下面数
				r := 0
				if dp[i][j+1]-dungeon[i][j] > 0 {
					r = dp[i][j+1] - dungeon[i][j]
				} else {
					r = 1
				}

				d := 0
				if dp[i+1][j]-dungeon[i][j] > 0 {
					d = dp[i+1][j] - dungeon[i][j]
				} else {
					d = 1
				}

				if r < d {
					dp[i][j] = r
				} else {
					dp[i][j] = d
				}
			}
		}

	}

	return dp[0][0]

}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	res := calculateMinimumHP([][]int{{-2, -3, 3}, {-5, -10, 1}, {10, 30, -5}})
	fmt.Print(res)
}
