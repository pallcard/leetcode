package main

import "sort"

// leetcode submit region begin(Prohibit modification and deletion)
func findMinArrowShots(points [][]int) int {
	//todo 先按照左边界排序
	// 如果当前气球的左边界大于上一个气球的右边界 左需要增加一根箭
	// 如果当前气球的左边界不大于上一个气球的右边界， 则更新当前气球的右边界为 上一个气球的右边界和当前气球右边界较小值
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	cnt := 1 //初始值为1，表示第一个气球已经用了一箭
	for i := 1; i < len(points); i++ {
		if points[i][0] > points[i-1][1] {
			cnt++
		} else if points[i][1] > points[i-1][1] {
			points[i][1] = points[i-1][1]
		}
	}
	return cnt
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
