package main

// leetcode submit region begin(Prohibit modification and deletion)
func combine(n int, k int) [][]int {

	res := make([][]int, 0)

	var backtrack func(cur int, trace []int)

	backtrack = func(cur int, trace []int) {
		if len(trace) == k {
			tempArr := make([]int, 0, k)
			tempArr = append(tempArr, trace...)
			res = append(res, tempArr)
			return
		}

		for i := cur; i <= n; i++ {
			trace = append(trace, i)
			backtrack(i+1, trace)
			trace = trace[:len(trace)-1]
		}
	}

	backtrack(1, []int{})
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
