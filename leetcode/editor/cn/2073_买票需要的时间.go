package main

// leetcode submit region begin(Prohibit modification and deletion)
func timeRequiredToBuy(tickets []int, k int) int {
	if len(tickets)-1 < k {
		return 0
	}

	res := 0
	for i, t := range tickets {
		if i <= k {
			res += min(t, tickets[k])
		} else {
			res += min(t, tickets[k]-1)
		}

	}
	return res
}

func min(x, y int) int {
	if x >= y {
		return y
	} else {
		return x
	}
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
