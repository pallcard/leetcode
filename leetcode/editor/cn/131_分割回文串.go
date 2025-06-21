package main

import "fmt"

// leetcode submit region begin(Prohibit modification and deletion)
func partition(s string) [][]string {
	//todo 回溯遍历出所有情况即可
	res := make([][]string, 0)

	var backtrack func(trace []string, index int)

	backtrack = func(trace []string, index int) {
		if index == len(s) {
			temp := make([]string, len(trace))
			copy(temp, trace)
			res = append(res, temp)
		}

		for i := index; i < len(s); i++ { //从index开始取，有index ～ (len(s)-1)种情况
			if isHuiwen(s[index : i+1]) { //以index开头的取1个，取2个，取3个...
				trace = append(trace, s[index:i+1])
				backtrack(trace, i+1)
				trace = trace[:len(trace)-1]
			}
		}
	}

	trace := make([]string, 0)
	backtrack(trace, 0)
	return res
}

func isHuiwen(s string) bool {
	if len(s) == 1 {
		return true
	}
	left := 0
	right := len(s) - 1

	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

	fmt.Print("xxx")
	i := partition("aab")
	fmt.Print(i)
}
