package main

import "fmt"

// leetcode submit region begin(Prohibit modification and deletion)
func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}

	if s == t {
		return s
	}

	left := 0
	right := 0

	window := map[byte]int{} //char:cnt
	need := map[byte]int{}   //char:cnt
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	valid := 0

	max := s + t

	for right < len(s) {
		char := s[right]
		right++
		window[char]++
		if window[char] == need[char] {
			valid++
			for valid == len(need) { //当窗口完全覆盖子串，开始缩小
				if len(max) > right-left {
					max = s[left:right]
				}

				lc := s[left]
				if window[lc] == need[lc] {
					valid--
				}
				window[lc]--
				left++
			}

		}
	}

	if len(max) > len(s) {
		return ""
	}

	return max
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	res := minWindow("a", "a")
	fmt.Print(res)
}
