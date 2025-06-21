package main

// leetcode submit region begin(Prohibit modification and deletion)
func longestPalindrome(s string) string {
	//todo 选择中心点(i), (i,i+1)向两边扩散
	// 注意回文串 可能是偶数，也可能是奇数
	if len(s) <= 0 {
		return ""
	}

	max := s[0:1]
	for i := 0; i < len(s)-1; i++ {
		palindrome1 := maxPalindrome(s, i, i)
		if len(palindrome1) > len(max) {
			max = palindrome1
		}
		palindrome2 := maxPalindrome(s, i, i+1)
		if len(palindrome2) > len(max) {
			max = palindrome2
		}

	}

	return max
}

func maxPalindrome(s string, left, right int) string {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return s[left+1 : right]
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
