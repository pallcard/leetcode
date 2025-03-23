package main

// leetcode submit region begin(Prohibit modification and deletion)
func canCompleteCircuit(gas []int, cost []int) int {
	// todo 计算每日剩余油数
	// 1. 总剩余油数 > 0
	// 2. 累计剩余油数 > 0,   累加时小于0，则从i+1开始累计， i+1的累计油数>0

	if len(gas) != len(cost) {
		return -1
	}

	sumGas := 0
	curGas := 0
	cur := -1
	for i := 0; i < len(gas); i++ {
		rest := gas[i] - cost[i]
		sumGas += rest
		curGas += rest
		if curGas < 0 {
			curGas = 0
			cur = -1
		} else if cur == -1 {
			cur = i
		}
	}
	if sumGas >= 0 && cur >= 0 {
		return cur
	}
	return -1
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
