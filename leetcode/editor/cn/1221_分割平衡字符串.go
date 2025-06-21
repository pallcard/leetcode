package main

// leetcode submit region begin(Prohibit modification and deletion)
func balancedStringSplit(s string) int {
	rCnt := 0
	lCnt := 0

	res := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'L' {
			lCnt++
			if lCnt == rCnt {
				res++
				lCnt = 0
				rCnt = 0
			}
		} else {
			rCnt++
			if lCnt == rCnt {
				res++
				lCnt = 0
				rCnt = 0
			}
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
