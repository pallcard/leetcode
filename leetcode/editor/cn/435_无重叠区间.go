package main

import (
	"fmt"
	"sort"
)

// leetcode submit region begin(Prohibit modification and deletion)
func eraseOverlapIntervals(intervals [][]int) int {

	//todo 按照又端点排序，
	// 当前节点左端点小于上一个节点又端点则需要删除当前节点，更新当前节点右端点用于下一次计算
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] <= intervals[j][1]
	})

	cnt := 0
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < intervals[i-1][1] {
			cnt++
			intervals[i][1] = intervals[i-1][1] //todo 删除了当前节点，故要更新当前节点右端点用于下一次计算
		}
	}
	return cnt
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	intervals := eraseOverlapIntervals([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}})
	fmt.Print(intervals)
}
