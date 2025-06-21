package main

import (
	"fmt"
	"strings"
)

// leetcode submit region begin(Prohibit modification and deletion)
func solveNQueens(n int) [][]string {

	res := make([][]string, 0)

	var backtrace func(trace []string)

	backtrace = func(trace []string) {
		if len(trace) == n {
			temp := make([]string, n, n)
			copy(temp, trace)
			res = append(res, temp)
			return
		}

		for i := 0; i < n; i++ {
			if validQueen(trace, i, n) {
				str := strings.Builder{}
				for j := 0; j < n; j++ {
					if i == j {
						str.Write([]byte("Q"))
					} else {
						str.Write([]byte("."))
					}
				}
				trace = append(trace, str.String())
				backtrace(trace)
				trace = trace[0 : len(trace)-1]
			}
		}

	}

	trace := make([]string, 0, n)
	backtrace(trace)
	return res
}

func validQueen(trace []string, line int, n int) bool {
	// todo 不同行，不同列，不同斜
	// 按照行递归，故不同行
	j := 1
	for i := len(trace) - 1; i >= 0; i-- {
		// 不同列
		if trace[i][line] == 'Q' {
			return false
		}

		//不同斜
		if ((line-j) >= 0 && trace[i][line-j] == 'Q') ||
			((line+j) < n && trace[i][line+j] == 'Q') {
			return false
		}
		j++
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	queens := solveNQueens(4)
	fmt.Print(queens)
}
