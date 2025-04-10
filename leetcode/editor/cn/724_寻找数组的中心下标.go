package main

// leetcode submit region begin(Prohibit modification and deletion)
func pivotIndex(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		return 0
	}
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	numSum := 0
	for i := 0; i < len(nums); i++ {
		// todo 注意 sum-nums[i]要为偶数
		if (sum-nums[i])%2 == 0 && (sum-nums[i])/2 == numSum {
			return i
		} else {
			numSum += nums[i]
		}
	}
	return -1
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
