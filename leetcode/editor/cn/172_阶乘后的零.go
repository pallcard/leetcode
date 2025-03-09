package main

import "fmt"

// leetcode submit region begin(Prohibit modification and deletion)
func trailingZeroes(n int) int {
	// todo n! 中质因子 2 的个数和质因子 5 的个数的较小值。
	cnt := 0
	for n > 0 {
		n /= 5
		cnt += n
	}
	return cnt
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

	var n int64 = 1
	var i int64 = 1
	for ; i <= 25; i++ {
		n *= i
	}
	fmt.Print(n)
}
