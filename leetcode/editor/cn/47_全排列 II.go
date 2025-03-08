package main

import (
	"fmt"
	"sort"
)

// leetcode submit region begin(Prohibit modification and deletion)
func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	var backtrack func(trace []int, used map[int]bool)
	backtrack = func(trace []int, used map[int]bool) {
		if len(trace) == len(nums) {
			tempArr := make([]int, 0, len(nums))
			tempArr = append(tempArr, trace...)
			res = append(res, tempArr)
			return
		}

		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			// todo 关键条件，前一个值和当前值相等，且前一个值未使用，则减枝
			if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}
			trace = append(trace, nums[i])
			used[i] = true
			backtrack(trace, used)
			trace = trace[:len(trace)-1]
			used[i] = false
		}

	}
	backtrack([]int{}, map[int]bool{})
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	unique := permuteUnique([]int{1, 1, 2})
	fmt.Print(unique)
}
