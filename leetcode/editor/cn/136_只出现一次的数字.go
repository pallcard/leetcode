package main

import "fmt"

// leetcode submit region begin(Prohibit modification and deletion)
func singleNumber(nums []int) int {
	if len(nums)%2 == 0 {
		return 0
	}

	res := nums[0]
	for i := 1; i < len(nums); i++ {
		res ^= nums[i]
	}

	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	solution := singleNumber([]int{1})
	fmt.Println(solution)
}
