package main

import (
	"fmt"
	"math"
)

// leetcode submit region begin(Prohibit modification and deletion)
func isRectangleCover(rectangles [][]int) bool {
	// x,y, a,b

	// 左下
	minX := math.MaxInt
	minY := math.MaxInt
	// 右上
	maxX := math.MinInt
	maxY := math.MinInt
	// todo 1.总面积相等
	// todo 2.顶点判断，当多个小矩形的顶点为奇数时才能为顶点
	/**
	 */
	pointSet := map[string]struct{}{}
	totalArea := 0
	for i := 0; i < len(rectangles); i++ {
		x1 := rectangles[i][0]
		y1 := rectangles[i][1]
		x2 := rectangles[i][2]
		y2 := rectangles[i][3]
		totalArea += (x2 - x1) * (y2 - y1)
		if minX > x1 {
			minX = x1
		}
		if minY > y1 {
			minY = y1
		}
		if maxX < x2 {
			maxX = x2
		}
		if maxY < y2 {
			maxY = y2
		}
		p1 := fmt.Sprintf("%d,%d", x1, y1)
		p2 := fmt.Sprintf("%d,%d", x1, y2)
		p3 := fmt.Sprintf("%d,%d", x2, y2)
		p4 := fmt.Sprintf("%d,%d", x2, y1)

		if _, ok := pointSet[p1]; ok {
			delete(pointSet, p1)
		} else {
			pointSet[p1] = struct{}{}
		}
		if _, ok := pointSet[p2]; ok {
			delete(pointSet, p2)
		} else {
			pointSet[p2] = struct{}{}
		}
		if _, ok := pointSet[p3]; ok {
			delete(pointSet, p3)
		} else {
			pointSet[p3] = struct{}{}
		}
		if _, ok := pointSet[p4]; ok {
			delete(pointSet, p4)
		} else {
			pointSet[p4] = struct{}{}
		}
	}

	// 面积相等
	if totalArea != (maxX-minX)*(maxY-minY) {
		return false
	}

	// 顶点相等
	if len(pointSet) != 4 {
		return false
	}
	maxP1 := fmt.Sprintf("%d,%d", minX, minY)
	maxP2 := fmt.Sprintf("%d,%d", minX, maxY)
	maxP3 := fmt.Sprintf("%d,%d", maxX, maxY)
	maxP4 := fmt.Sprintf("%d,%d", maxX, minY)
	if _, ok := pointSet[maxP1]; !ok {
		return false
	}
	if _, ok := pointSet[maxP2]; !ok {
		return false
	}
	if _, ok := pointSet[maxP3]; !ok {
		return false
	}
	if _, ok := pointSet[maxP4]; !ok {
		return false
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
