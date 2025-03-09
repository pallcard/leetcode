package main

import "fmt"

// leetcode submit region begin(Prohibit modification and deletion)
func hammingWeight(n int) int {
	if n == 0 {
		return 0
	}
	cnt := 1
	for n&(n-1) != 0 {
		n &= n - 1
		cnt++

	}
	return cnt
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	hammingWeight(11)
	fmt.Print(byte(11))
}
