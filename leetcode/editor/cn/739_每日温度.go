package main

import "fmt"

//leetcode submit region begin(Prohibit modification and deletion)

func dailyTemperatures(temperatures []int) []int {
	//todo 单调栈
	// 求右边比左边大的元素， 单调递减的栈

	res := make([]int, len(temperatures), len(temperatures))

	stack := make([]int, 0, len(temperatures))

	for i := 0; i < len(temperatures); i++ {
		// todo 栈空或当前元素<=栈顶， 入栈
		if len(stack) <= 0 || temperatures[stack[len(stack)-1]] >= temperatures[i] {
			stack = append(stack, i)
		} else {
			//todo 栈不空 && 当前元素 > 栈顶， 出栈
			for len(stack) > 0 && temperatures[stack[len(stack)-1]] < temperatures[i] {
				top := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				res[top] = i - top
			}
			// todo 出完之后在把当前元素入栈
			stack = append(stack, i)
		}
	}

	// todo 处理栈中剩余元素
	for len(stack) > 0 {
		res[stack[len(stack)-1]] = 0
		stack = stack[:len(stack)-1]
	}

	return res

}

func dailyTemperatures2(temperatures []int) []int {

	res := make([]int, 0)
	for i := 0; i < len(temperatures); i++ {

		index := 0
		high := false
		for j := i + 1; j < len(temperatures); j++ {
			index++
			if temperatures[j] > temperatures[i] {
				high = true
				break
			}
		}
		if high {
			res = append(res, index)
		} else {
			res = append(res, 0)
		}

	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	temperatures := dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73})
	fmt.Print(temperatures)
}
