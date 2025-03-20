package main

// leetcode submit region begin(Prohibit modification and deletion)
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	}
	max := nums[0]
	pre := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] >= 0 {
			if pre > 0 {
				pre += nums[i]
			} else {
				pre = nums[i]
			}

		} else { //nums[i]<0
			if pre > 0 {
				pre += nums[i]
			} else {
				pre = nums[i]

			}
		}
		if max < pre {
			max = pre
		}

	}
	return max
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
