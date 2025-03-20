package main

// leetcode submit region begin(Prohibit modification and deletion)
func wiggleMaxLength(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	max := 1
	pre := nums[0]
	dir := 0 // 1增 -1减
	for i := 1; i < len(nums); i++ {
		if dir == 0 {
			if nums[i] > pre {
				pre = nums[i]
				dir = 1
				max++
			} else if nums[i] < pre {
				pre = nums[i]
				dir = -1
				max++
			}
		} else {
			if dir == 1 {
				if nums[i] < pre {
					pre = nums[i]
					dir = -1
					max++
				} else if nums[i] >= pre {
					pre = nums[i]
				}
			} else {
				if nums[i] > pre {
					pre = nums[i]
					dir = 1
					max++
				} else if nums[i] <= pre {
					pre = nums[i]
				}
			}
		}
	}
	return max
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
