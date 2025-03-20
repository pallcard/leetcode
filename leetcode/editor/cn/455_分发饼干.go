package main

import "sort"

// leetcode submit region begin(Prohibit modification and deletion)
func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	sIndex := 0
	gIndex := 0
	cnt := 0
	for sIndex < len(s) && gIndex < len(g) {
		if s[sIndex] >= g[gIndex] {
			cnt++
			sIndex++
			gIndex++
		} else if s[sIndex] < g[gIndex] {
			sIndex++
		}
	}
	return cnt
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
