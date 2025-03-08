package main

import (
	"fmt"
	"math"
)

// leetcode submit region begin(Prohibit modification and deletion)
func minDistance(word1 string, word2 string) int {
	// todo dp定义为 word1[0...i] word2[0...j]的最小编辑距离是 dp[i+1][j+1]
	/**
	w1\w2  ""  a  p  p  l  e
	""     0   1  2  3  4  5
	r      1
	a      2
	d      3               ?
	**/

	dp := make([][]int, len(word1)+1)
	for i := 0; i < len(word1)+1; i++ {
		dp[i] = make([]int, len(word2)+1)
		dp[i][0] = i
	}

	for j := 0; j < len(word2)+1; j++ {
		dp[0][j] = j
	}

	for i := 1; i < len(word1)+1; i++ {
		for j := 1; j < len(word2)+1; j++ {
			if word1[i-1] == word2[j-1] { //todo 下标
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = minNum(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
			}
		}
	}

	return dp[len(word1)][len(word2)]
}

func minDistance2(word1 string, word2 string) int {
	// todo 两个字符穿一般定义两个指针分别指向字符串
	memo := make([][]int, 0, len(word1))
	for i := 0; i < len(word1); i++ {
		memo = append(memo, make([]int, len(word2)))
		for j := 0; j < len(word2); j++ {
			memo[i][j] = -1
		}
	}

	return minDistanceDp(word1, len(word1)-1, word2, len(word2)-1, memo)
}

func minDistanceDp(word1 string, i int, word2 string, j int, memo [][]int) int {
	// base case
	if i == -1 {
		return j + 1
	}
	if j == -1 {
		return i + 1
	}

	if memo[i][j] != -1 {
		return memo[i][j]
	}

	if word1[i] == word2[j] {
		memo[i][j] = minDistanceDp(word1, i-1, word2, j-1, memo)
	} else {
		// word1删除
		a := minDistanceDp(word1, i-1, word2, j, memo)
		// word1插入
		b := minDistanceDp(word1, i, word2, j-1, memo)
		// word1替换
		c := minDistanceDp(word1, i-1, word2, j-1, memo)
		memo[i][j] = minNum(a, b, c) + 1
	}
	return memo[i][j]
}

func minNum(nums ...int) int {
	min := math.MaxInt
	for _, i := range nums {
		if min > i {
			min = i
		}
	}
	return min
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	distance := minDistance("abc", "a")
	fmt.Print(distance)
}
