package main

import "sort"

// leetcode submit region begin(Prohibit modification and deletion)
func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	res := make([][]int, 0)

	var backtrace1 func(target, curIndex int, trace []int)
	backtrace1 = func(target, curIndex int, trace []int) {
		if target == 0 {
			temp := make([]int, 0, len(trace))
			temp = append(temp, trace...)
			res = append(res, temp)
		}

		if target < 0 {
			return
		}

		for i := curIndex; i < len(candidates); i++ {
			trace = append(trace, candidates[i])
			backtrace1(target-candidates[i], i, trace)
			trace = trace[:len(trace)-1]
		}

	}

	trace := make([]int, 0)
	backtrace1(target, 0, trace)

	return res

}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	combinationSum([]int{2, 3, 6, 7}, 7)
}
