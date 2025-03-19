package main

// leetcode submit region begin(Prohibit modification and deletion)
func largestRectangleArea(heights []int) int {
	if len(heights) == 0 {
		return 0
	}

	//todo 以i为高度的最大矩形
	// height[i]*(minRightHeight_i -minRightHeight_i -1)
	// 寻找以i为中心的左边第一个小于i的高度的坐标
	// 寻找以i为中心的右边第一个大于i的高度的坐标

	leftH := make([]int, len(heights), len(heights))
	rightH := make([]int, len(heights), len(heights))

	leftH[0] = -1
	for i := 1; i < len(heights); i++ { //todo 下标需要特别注意
		t := i - 1
		for t >= 0 && heights[i] <= heights[t] { //todo 关键步骤
			t = leftH[t]
		}
		leftH[i] = t
	}

	rightH[len(heights)-1] = len(heights)
	for i := len(heights) - 2; i >= 0; i-- {
		t := i + 1
		for t < len(heights) && heights[i] <= heights[t] {
			t = rightH[t]
		}
		rightH[i] = t
	}

	maxArea := 0
	for i := 0; i < len(heights); i++ {
		temp := heights[i] * (rightH[i] - leftH[i] - 1)
		if temp > maxArea {
			maxArea = temp
		}
	}
	return maxArea
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
