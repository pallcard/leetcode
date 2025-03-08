package main

import "fmt"

// leetcode submit region begin(Prohibit modification and deletion)
func combinationSum3(k int, n int) [][]int {
	if k > 9 || n > 45 {
		return [][]int{}
	}

	used := map[int]bool{}

	res := make([][]int, 0)
	var backtrace3 func(k, n, cur int, used map[int]bool, trace []int)
	backtrace3 = func(k, n, cur int, used map[int]bool, trace []int) {
		if k == 0 && n == 0 {
			temp := make([]int, 0, len(trace))
			temp = append(temp, trace...)
			res = append(res, temp)
			return
		}

		if k < 0 || n < 0 {
			return
		}

		for i := cur+1; i <= 9; i++ {
			if used[i] {
				continue
			}
			used[i] = true
			trace = append(trace, i)
			backtrace3(k-1, n-i, i, used, trace)
			used[i] = false
			trace = trace[:len(trace)-1]
		}
	}

	trace := make([]int, 0)
	backtrace3(k, n, 0, used, trace)
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	sum3 := combinationSum3(3, 7)
	fmt.Print(sum3)
}
