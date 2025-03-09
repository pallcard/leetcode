package main

// leetcode submit region begin(Prohibit modification and deletion)
func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	dp := make([]int, len(nums))

	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if i == 1 {
			dp[i] = max(dp[i-1], nums[i])
		} else {
			dp[i] = max(dp[i-1], nums[i]+dp[i-2]) // 不偷这家， 偷这家
		}
	}
	return dp[len(nums)-1]
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
