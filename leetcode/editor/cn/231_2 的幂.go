package main

// leetcode submit region begin(Prohibit modification and deletion)
func isPowerOfTwo(n int) bool {
	if n == 0 {
		return false
	}
	if n&(n-1) != 0 {
		return false
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
