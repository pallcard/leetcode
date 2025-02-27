package main

import "fmt"

// leetcode submit region begin(Prohibit modification and deletion)
func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	left := 0
	right := 0

	window := map[byte]int{} //char : cnt
	need := map[byte]int{}   //char : cnt
	for _, c := range s1 {
		need[byte(c)]++
	}

	valid := 0

	for ; right < len(s2); right++ {
		char := s2[right]
		if _, ok := need[char]; ok {
			window[char]++
			if window[char] == need[char] {
				valid++
				if valid == len(need) {
					return true
				}
			} else if window[char] > need[char] {
				// 缩小窗口到char的位置
				for s2[left] != char {
					lc := s2[left]
					if window[lc] == need[lc] {
						valid--
					}
					window[lc]--
					left++
				}

			}
		} else { //出现不存在的，窗口直接重置到 right+1 处
			window = map[byte]int{}
			valid = 0
			left = right + 1
		}

	}
	return false

}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	res := checkInclusion("ab", "eidboaoo")
	fmt.Println(res)
}
