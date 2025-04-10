package main

// leetcode submit region begin(Prohibit modification and deletion)
func uniqueOccurrences(arr []int) bool {
	numCntMap := map[int]int{}

	for i := 0; i < len(arr); i++ {
		numCntMap[arr[i]]++
	}

	repeatSet := map[int]struct{}{}

	for _, v := range numCntMap {
		if _, ok := repeatSet[v]; ok {
			return false
		}
		repeatSet[v] = struct{}{}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	uniqueOccurrences([]int{1, 2})
}
