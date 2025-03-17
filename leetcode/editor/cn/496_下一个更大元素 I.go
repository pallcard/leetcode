package main

import "fmt"

// leetcode submit region begin(Prohibit modification and deletion)
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	stack := make([]int, 0, len(nums2))

	resMap := map[int]int{}
	for i := 0; i < len(nums2); i++ {
		//todo 构造递减的单调栈
		if len(stack) <= 0 || nums2[stack[len(stack)-1]] >= nums2[i] {
			stack = append(stack, i)
		} else {
			for len(stack) > 0 && nums2[stack[len(stack)-1]] < nums2[i] {

				top := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				resMap[nums2[top]] = nums2[i]
			}
			stack = append(stack, i)

		}

	}

	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		resMap[nums2[top]] = -1
	}

	res := make([]int, 0, len(nums1))
	for i := 0; i < len(nums1); i++ {
		res = append(res, resMap[nums1[i]])
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	element := nextGreaterElement([]int{4, 1, 2}, []int{1, 3, 4, 2})
	fmt.Print(element)
}
