package main

import "sort"

// leetcode submit region begin(Prohibit modification and deletion)
func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] < intervals[j][0] {
			return true
		} else if intervals[i][0] > intervals[j][0] {
			return false
		}
		return intervals[i][1] < intervals[j][1]
	})

	res := make([][]int, 0, len(intervals))
	left := intervals[0][0]
	right := intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= right {
			right = max(intervals[i][1], right)
		} else {
			res = append(res, []int{left, right})
			left = intervals[i][0]
			right = intervals[i][1]
		}
	}
	res = append(res, []int{left, right})
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
