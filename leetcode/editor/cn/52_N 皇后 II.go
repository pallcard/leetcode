package main

// leetcode submit region begin(Prohibit modification and deletion)
func totalNQueens(n int) int {
	res := 0

	var backtrack func(trace [][]byte)

	backtrack = func(trace [][]byte) {
		if len(trace) == n {
			res++
			return
		}

		for i := 0; i < n; i++ {
			if validQueen2(trace, i, n) {
				temp := make([]byte, n)
				temp[i] = 1
				trace = append(trace, temp)
				backtrack(trace)
				trace = trace[:len(trace)-1]
			}

		}

	}

	trace := make([][]byte, 0, n)
	backtrack(trace)

	return res
}

func validQueen2(trace [][]byte, point, n int) bool {
	j := 1
	for i := len(trace) - 1; i >= 0; i-- {
		if trace[i][point] == 1 {
			return false
		}

		if (point-j >= 0 && trace[i][point-j] == 1) ||
			(point+j < n && trace[i][point+j] == 1) {
			return false
		}
		j++
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	totalNQueens(4)
}
