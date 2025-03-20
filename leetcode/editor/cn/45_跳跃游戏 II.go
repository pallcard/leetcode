package main

import "fmt"

// leetcode submit region begin(Prohibit modification and deletion)
func jump(nums []int) int {

	res := make([]int, len(nums), len(nums))

	for i := len(nums) - 2; i >= 0; i-- {
		minStep := len(nums)
		for j := 0; j < nums[i]; j++ {
			if i+j+1 < len(res) && minStep > res[i+j+1]+1 {
				minStep = res[i+j+1] + 1
			}
		}
		res[i] = minStep
	}
	if res[0] >= len(nums) { //不可达
		return -1
	}
	return res[0]
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	res := jump([]int{2, 1})
	fmt.Print(res)
}
