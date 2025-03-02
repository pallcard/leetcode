package main

// leetcode submit region begin(Prohibit modification and deletion)
func fib(n int) int {
	if n <= 1 {
		return n
	}
	memo := make([]int, n+1)
	return fibMemo(n, memo) //fib(n-1) + fib(n-2)
}

func fibMemo(n int, memo []int) int {
	if n <= 1 {
		return n
	}

	if memo[n] > 0 {
		return memo[n]
	}

	memo[n] = fib(n-1) + fib(n-2)
	return memo[n]
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
