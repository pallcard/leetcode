package main

import "fmt"

// leetcode submit region begin(Prohibit modification and deletion)
func trap(height []int) int {
	//todo left[i] 第i个数字左边最大的数
	//todo right[i] 第i个数字右边最大的数

	if len(height) <= 1 {
		return 0
	}

	leftH := make([]int, len(height), len(height))
	rightH := make([]int, len(height), len(height))

	leftH[0] = height[0]
	for i := 1; i < len(height); i++ {
		if height[i] >= leftH[i-1] {
			leftH[i] = height[i]
		} else {
			leftH[i] = leftH[i-1]
		}
	}

	rightH[len(height)-1] = height[len(height)-1]
	for i := len(height) - 2; i >= 0; i-- {
		if height[i] >= rightH[i+1] {
			rightH[i] = height[i]
		} else {
			rightH[i] = rightH[i+1]
		}
	}

	water := 0
	for i := 1; i < len(height); i++ {
		if height[i] < leftH[i] && height[i] < rightH[i] {
			if leftH[i] < rightH[i] {
				water += leftH[i] - height[i]
			} else {
				water += rightH[i] - height[i]
			}
		}
	}

	return water
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	res := trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})
	fmt.Print(res)
}
