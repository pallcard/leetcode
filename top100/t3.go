package top100

// tmmzuxt
func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}

	maxLen := 1

	left := 0
	right := 0

	charSet := map[uint8]int{} // char:index

	for right <= len(s)-1 {
		// index >= left 保证在窗口内，优化for循环剔除数据
		if index, ok := charSet[s[right]]; ok && index >= left {
			left = index + 1
			charSet[s[right]] = right
			right++
		} else {
			curLen := right - left + 1
			if maxLen < curLen {
				maxLen = curLen
			}
			charSet[s[right]] = right
			right++
		}
	}
	return maxLen
}
