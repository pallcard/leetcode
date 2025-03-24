package main

import "fmt"

// leetcode submit region begin(Prohibit modification and deletion)
func partitionLabels(s string) []int {
	// todo 只需要存end即可
	charMap := map[uint8][2]int{} // char:[start, end]
	for i := 0; i < len(s); i++ {
		if m, ok := charMap[s[i]]; ok {
			m[1] = i
			charMap[s[i]] = m
		} else {
			charMap[s[i]] = [2]int{i, i}
		}
	}

	res := make([]int, 0)
	left := 0
	right := 0
	for i := 0; i < len(s); i++ {
		// todo 如果当前下标大于right，表示到了新边界
		if i > right {
			res = append(res, right-left+1)
			left = i
			right = i
		}
		// todo 当区间类出现了比当前边界大的数，则更新当前边界
		if charMap[s[i]][1] > right {
			right = charMap[s[i]][1]
		}
	}
	res = append(res, right-left+1)

	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	labels := partitionLabels("ababcbacadefegdehijhklij")
	fmt.Print(labels)
}
