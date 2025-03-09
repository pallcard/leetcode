package main

// leetcode submit region begin(Prohibit modification and deletion)
func countPrimes(n int) int {
	arr := make([]bool, n)
	for i := 2; i < len(arr); i++ {
		arr[i] = true
	}

	for i := 2; i < n; i++ {
		//for j := i; i*j < n; j++ {
		//	arr[i*j] = false
		//}
		for j := i * i; j < n; j += i {
			arr[j] = false
		}
	}
	cnt := 0
	for i := 2; i < n; i++ {
		if arr[i] {
			cnt++
		}
	}

	return cnt
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
