package main

import "fmt"

// leetcode submit region begin(Prohibit modification and deletion)
func canPartition(nums []int) bool {
	if len(nums) <= 1 {
		return false
	}
	//todo
	target := 0
	for _, num := range nums {
		target += num
	}
	if target%2 != 0 {
		return false
	}

	target /= 2
	find := false
	var backtrack func(target, curIndex int)
	backtrack = func(target, curIndex int) {
		if find {
			return
		}
		if target == 0 {
			find = true
			return
		} else if target < 0 || curIndex >= len(nums) {
			return
		}

		for i := curIndex; i < len(nums); i++ {
			backtrack(target-nums[i], i+1)
		}

	}

	backtrack(target, 0)
	return find
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	b := canPartition([]int{1, 5, 11, 5})
	fmt.Print(b)
}
