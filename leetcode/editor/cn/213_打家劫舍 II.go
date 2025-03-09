package main

// leetcode submit region begin(Prohibit modification and deletion)
func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	// todo 增加第1家和最后一家不能同时抢，将环形拆分成两个，[1 , N-1] 和 [2, N]两种情况，取最大值即可
	return max(robRang(nums, 0, len(nums)-2), robRang(nums, 1, len(nums)-1))
}

func robRang(nums []int, start int, end int) int {
	//dp := make([]int, len(nums))
	//dp[start] = nums[start]
	//dp[start+1] = max(nums[start], nums[start+1])
	dpCurPre := 0        //当前值的前一个
	dpCur := nums[start] // 当前值 todo 把数组压缩成变量
	for i := start + 1; i <= end; i++ {
		//dp[i] = max(dp[i-1], dp[(i-2]+nums[i])
		//pre := dpCurPre+nums[i]
		//dpCurPre = dpCur
		//dpCur = max(dpCur, pre)
		temp := max(dpCur, dpCurPre+nums[i])
		dpCurPre = dpCur
		dpCur = temp
	}
	return dpCur
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
