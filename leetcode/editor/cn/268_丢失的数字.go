package main

// leetcode submit region begin(Prohibit modification and deletion)
func missingNumber(nums []int) int {
	num := 0
	for i := 0; i < len(nums); i++ {
		num ^= nums[i]
		num ^= i + 1
	}

	return num
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
