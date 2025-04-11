package main

// leetcode submit region begin(Prohibit modification and deletion)
func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	dict := map[byte]byte{}

	repeatT := map[byte]bool{} //todo 还要排除重复映射的
	// abcd
	// abbb

	for i := 0; i < len(s); i++ {
		if c, ok := dict[s[i]]; ok {
			if t[i] != c {
				return false
			}
		} else {
			if repeatT[t[i]] {
				return false
			}
			dict[s[i]] = t[i]
			repeatT[t[i]] = true
		}
	}

	return true

}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
