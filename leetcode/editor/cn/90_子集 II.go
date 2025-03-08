package main

import "sort"

// leetcode submit region begin(Prohibit modification and deletion)
func subsetsWithDup(nums []int) [][]int {
	res := make([][]int, 0)

	sort.Ints(nums)

	var backtrack func(trace []int, size, cur int)

	backtrack = func(trace []int, size, cur int) {
		//if size == len(trace) {
			tempArr := make([]int, 0)
			tempArr = append(tempArr, trace...)
			res = append(res, tempArr)
		//}


		for i := cur; i < len(nums); i++ {

			// todo i>cur是关键,可以保证
			if i > cur && nums[i-1] == nums[i] {
				continue
			}
			trace = append(trace, nums[i])
			backtrack(trace, size, i+1)
			trace = trace[:len(trace)-1]

		}

	}

	//for i := 0; i <= len(nums); i++ {
		backtrack([]int{}, 0, 0)
	//}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
