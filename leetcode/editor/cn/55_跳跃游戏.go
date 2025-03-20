package main

// leetcode submit region begin(Prohibit modification and deletion)
func canJump(nums []int) bool {
	if len(nums) == 0 {
		return false
	}
	if len(nums) == 1 && nums[0] == 0 {
		return true
	}
	need := 1
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] < need {
			need++
		} else {
			need = 1
		}
	}

	return need == 1

}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
