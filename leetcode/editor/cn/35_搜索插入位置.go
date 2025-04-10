package main

// leetcode submit region begin(Prohibit modification and deletion)
func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			return mid
		}
	}
	return left
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	searchInsert([]int{1, 2, 3, 5, 6}, 4)
}
