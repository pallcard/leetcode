package main

// leetcode submit region begin(Prohibit modification and deletion)
func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	left := 0
	right := 0

	// 定义窗口
	window := map[byte]int{}
	max := 0
	for ; right < len(s); right++ {
		_, ok := window[s[right]]
		if !ok {
			window[s[right]] = right
		} else {
			if len(window) > max {
				max = len(window)
			}
			// 窗口缩小
			for s[left] != s[right] {
				delete(window, s[left])
				left++
			}
			left++
		}
	}
	if len(window) > max {
		max = len(window)
	}

	return max
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
