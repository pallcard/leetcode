package main

import "fmt"

// leetcode submit region begin(Prohibit modification and deletion)
func findTargetSumWays(nums []int, target int) int {

	cnt := 0
	var backtrack func(target int, curIndex int)
	backtrack = func(target int, curIndex int) {
		if target == 0 && curIndex == len(nums) {
			cnt++
		}
		if curIndex >= len(nums) {
			return
		}
		backtrack(target+nums[curIndex], curIndex+1)
		backtrack(target-nums[curIndex], curIndex+1)
	}

	backtrack(-target, 0)

	return cnt
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	cnt := findTargetSumWays([]int{1, 1, 1, 1, 1}, -3)
	fmt.Print(cnt)
}
