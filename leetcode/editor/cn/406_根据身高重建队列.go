package main

import (
	"fmt"
	"sort"
)

// leetcode submit region begin(Prohibit modification and deletion)
func reconstructQueue(people [][]int) [][]int {
	// 两个纬度，先按照一个纬度排序，在处理另外一个纬度
	sort.Slice(people, func(i, j int) bool {
		// [0]倒叙 [1]正序
		if people[i][0] > people[j][0] {
			return true
		} else if people[i][0] < people[j][0] {
			return false
		} else {
			return people[i][1] < people[j][1]
		}
	})

	for i := 0; i < len(people); i++ {
		if i == people[i][1] {
			continue
		} else {
			// 将people[i]从people[i][1]移动到i的位置
			src := people[i][1]
			temp := people[i]
			copy(people[src+1:i+1], people[src:i])
			people[src] = temp
		}
	}

	return people

}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	queue := reconstructQueue([][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}})
	fmt.Print(queue)
}
