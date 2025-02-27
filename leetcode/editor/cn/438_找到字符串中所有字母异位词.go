package main

import "fmt"

// leetcode submit region begin(Prohibit modification and deletion)
func findAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return []int{}
	}

	if len(s) == 0 {
		return []int{0}
	}

	left := 0
	right := 0

	window := map[byte]int{} //char:cnt
	need := map[byte]int{}   //char:cnt
	for _, char := range p {
		need[byte(char)]++
	}

	valid := 0

	res := make([]int, 0)

	for ; right < len(s); right++ {
		char := s[right]

		if _, ok := need[char]; ok {
			window[char]++
			if window[char] == need[char] {
				valid++
				if valid == len(need) {
					res = append(res, left)
				}
			} else if window[char] > need[char] { // 缩小窗口
				for s[left] != char {
					leftChar := s[left]
					if window[leftChar] == need[leftChar] {
						valid--
					}
					window[leftChar]--
					left++
				}
				window[s[left]]--
				left++
				if valid == len(need) {
					res = append(res, left)
				}
			}
		} else { // 出现其他字符
			window = map[byte]int{}
			left = right + 1
			valid = 0
		}
	}

	return res

}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	res := findAnagrams("cbaebabacd", "acb")
	fmt.Print(res)
}
