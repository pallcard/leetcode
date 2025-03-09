package main

// leetcode submit region begin(Prohibit modification and deletion)
func canWinNim(n int) bool {
	//todo 每次能拿1-3个石头，最后保证4个石头的时候给到对方，对方就输了
	// 谁踩着4的倍数谁输
	return n%4 != 0
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
