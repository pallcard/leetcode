package main

// leetcode submit region begin(Prohibit modification and deletion)
func lemonadeChange(bills []int) bool {
	coin5 := 0
	coin10 := 0

	for i := 0; i < len(bills); i++ {
		switch bills[i] {
		case 5:
			coin5++
		case 10:
			if coin5 > 0 {
				coin5--
				coin10++
			} else {
				return false
			}
		case 20:
			rest := 15
			if coin10 > 0 {
				coin10--
				rest -= 10
			}
			coin5 -= rest / 5
			if coin5 < 0 {
				return false
			}

		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
