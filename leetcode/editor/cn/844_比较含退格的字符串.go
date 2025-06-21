package main

// leetcode submit region begin(Prohibit modification and deletion)
func backspaceCompare(s string, t string) bool {
	// todo 边遍历边比较
	// todo 也可以先把字符串处理成最终字符串
	sIndex := len(s) - 1
	tIndex := len(t) - 1

	for sIndex >= 0 || tIndex >= 0 { //todo 使用或
		cnt := 0
		for sIndex >= 0 && (s[sIndex] == '#' || cnt > 0) {
			if s[sIndex] == '#' {
				sIndex--
				cnt++
			} else {
				cnt--
				sIndex--
			}

		}

		cnt = 0
		for tIndex >= 0 && (t[tIndex] == '#' || cnt > 0) {
			if t[tIndex] == '#' {
				tIndex--
				cnt++
			} else {
				cnt--
				tIndex--
			}
		}

		if sIndex >= 0 && tIndex >= 0 {
			if s[sIndex] != t[tIndex] {
				return false
			}
			sIndex--
			tIndex--
		} else if sIndex >= 0 || tIndex >= 0 {
			return false
		} else {
			return true
		}

	}

	//cnt := 0
	//for sIndex >= 0 && (s[sIndex] == '#' || cnt > 0) {
	//	if s[sIndex] == '#' {
	//		sIndex--
	//		cnt++
	//	} else {
	//		cnt--
	//		sIndex--
	//	}
	//
	//}
	//
	//cnt = 0
	//for tIndex >= 0 && (t[tIndex] == '#' || cnt > 0) {
	//	if t[tIndex] == '#' {
	//		tIndex--
	//		cnt++
	//	} else {
	//		cnt--
	//		tIndex--
	//	}
	//}

	return true

}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	backspaceCompare("nzp#o#g", "b#nzp#o#g")
}
