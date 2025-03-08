package main

import "sort"

// leetcode submit region begin(Prohibit modification and deletion)
func combinationSum2(candidates []int, target int) [][]int {
	// 1，7，1^，
	sort.Ints(candidates) //排序，去掉 1 7 ｜ 7 1 的情况
	res := make([][]int, 0)

	var backtrace2 func(target int, curIndex int, trace []int)

	backtrace2 = func(target int, curIndex int, trace []int) {
		if target == 0 {
			temp := make([]int, 0, len(trace))
			temp = append(temp, trace...)
			res = append(res, temp)
		}

		if target < 0 {
			return
		}

		//useMap := map[int]bool{}
		for i := curIndex; i < len(candidates); i++ {
			//if useMap[candidates[i]] {
			//	continue
			//} else {
			//	useMap[candidates[i]] = true
			//}
			if i > curIndex && candidates[i] == candidates[i-1] { //去掉 1 7 ｜ 1^ 7的情况
				continue
			}
			trace = append(trace, candidates[i])
			backtrace2(target-candidates[i], i+1, trace)
			trace = trace[:len(trace)-1]
		}

	}

	backtrace2(target, 0, []int{})
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
