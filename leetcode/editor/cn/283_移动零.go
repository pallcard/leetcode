package main

// leetcode submit region begin(Prohibit modification and deletion)
func moveZeroes(nums []int) {
	if len(nums) == 0 {
		return
	}

	left := 0
	for left < len(nums) && nums[left] != 0 {
		left++
	}

	for i := left+1; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[left] = nums[left],nums[i]
			left++
		}
	}

	//rightIndex := len(nums) - 1
	//
	//for rightIndex >= 0 && nums[rightIndex] == 0 {
	//	rightIndex--
	//}
	//
	//for i := 0; i < rightIndex; i++ {
	//	if nums[i] == 0 {
	//		nums[i], nums[rightIndex] = nums[rightIndex],nums[i]
	//		rightIndex--
	//	}
	//}
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
