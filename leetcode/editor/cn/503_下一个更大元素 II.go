package main

import "fmt"

// leetcode submit region begin(Prohibit modification and deletion)
func nextGreaterElements(nums []int) []int {
	nums2 := make([]int, 0, len(nums)*2)
	nums2 = append(nums2, nums...)
	nums2 = append(nums2, nums...)

	stack := make([]int, 0)
	res := make([]int, len(nums), len(nums))
	for i := 0; i < len(res); i++ {
		res[i] = -1
	}
	for i := 0; i < len(nums2); i++ {
		if len(stack) <= 0 || nums2[stack[len(stack)-1]] >= nums2[i] {
			stack = append(stack, i)
		} else {
			//todo for循环遍历
			for len(stack) > 0 && nums2[stack[len(stack)-1]] < nums2[i] {
				top := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if top < len(nums) {
					res[top] = nums2[i]
				}
			}

			stack = append(stack, i)
		}
	}

	return res
}

func nextGreaterElements2(nums []int) []int {
	// 拼接一个新的nums
	numsNew := make([]int, len(nums)*2)
	copy(numsNew, nums)
	copy(numsNew[len(nums):], nums)
	// 用新的nums大小来初始化result
	result := make([]int, len(numsNew))
	for i := range result {
		result[i] = -1
	}

	// 开始单调栈
	st := []int{0}
	for i := 1; i < len(numsNew); i++ {
		if numsNew[i] < numsNew[st[len(st)-1]] {
			st = append(st, i)
		} else if numsNew[i] == numsNew[st[len(st)-1]] {
			st = append(st, i)
		} else {
			for len(st) > 0 && numsNew[i] > numsNew[st[len(st)-1]] {
				result[st[len(st)-1]] = numsNew[i]
				st = st[:len(st)-1]
			}
			st = append(st, i)
		}
	}
	result = result[:len(result)/2]
	return result
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	res := nextGreaterElements([]int{1, 2, 1})
	fmt.Print(res)
}
