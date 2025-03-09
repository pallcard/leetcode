package main

import "fmt"

// leetcode submit region begin(Prohibit modification and deletion)
func predictTheWinner(nums []int) bool {
	// todo dp[i][j][k] k取0 (先手)、1 (后手), i、j表示当前数字是[i,j]
	// todo dp[i][j][0] 当前数字[i,j]时先手的最佳得分，
	// todo dp[i][j][1] 当前数字[i,j]时后手的最佳得分，
	// todo dp[i][i][0]=nums[i], 只有一个数字时，先手的最佳得分就是这个数字
	// todo dp[i][i][1]=0,只有一个数字时，后手的最佳得分就是0
	// todo dp[0][len(nums)-1][1] - dp[0][len(nums)-1][0] 先手的得分与后手的得分比较大小

	if len(nums) == 1 {
		return true
	}

	if len(nums)%2 == 0 { //偶数堆必赢，可以先看把数字按照下标的奇偶分成两堆，如果下标奇数堆数值大，则每次选奇数堆
		return true
	}

	dp := make([][][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = make([][]int, len(nums))
		for j := 0; j < len(nums); j++ {
			dp[i][j] = make([]int, 2)
		}
	}

	for i := 0; i < len(nums); i++ {
		dp[i][i][0] = nums[i] //仅仅一个数字时，先手取
		dp[i][i][1] = 0
	}

	for i := len(nums) - 2; i >= 0; i-- {
		for j := i + 1; j < len(nums); j++ {
			left := dp[i+1][j][1] + nums[i]  //取左边的
			right := dp[i][j-1][1] + nums[j] //取右边的
			if left > right {                //左边>右边
				dp[i][j][0] = left          // 先手取左边
				dp[i][j][1] = dp[i+1][j][0] //后手则取上一次的先手的值
			} else { // 左边<=右边
				dp[i][j][0] = right
				dp[i][j][1] = dp[i][j-1][0]
			}

		}
	}

	return dp[0][len(nums)-1][0] >= dp[0][len(nums)-1][1]

}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	res := predictTheWinner([]int{1, 1})
	fmt.Print(res)
}
