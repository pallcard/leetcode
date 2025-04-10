package main

// leetcode submit region begin(Prohibit modification and deletion)
func rotate(nums []int, k int) {
	//  1 2 3 4 5 6 7
	//  4 3 2 1 7 6 5
	//  5 6 7 1 2 3 4

	// todo 注意小细节，边界条件
	if len(nums) <= 1 {
		return
	}
	k %= len(nums) //todo 注意k大于数组长度

	swap(nums, 0, len(nums)-1-k)
	swap(nums, len(nums)-k, len(nums)-1)
	swap(nums, 0, len(nums)-1)

}

func swap(nums []int, left, right int) {
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
