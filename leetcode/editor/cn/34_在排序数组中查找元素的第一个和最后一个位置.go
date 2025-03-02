package main

import "fmt"

// leetcode submit region begin(Prohibit modification and deletion)
func searchRange(nums []int, target int) []int {

	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			minIndex := mid - 1
			for ; minIndex >= 0; minIndex-- {
				if nums[minIndex] != target {
					break
				}
			}
			minIndex++
			maxIndex := mid + 1
			for ; maxIndex < len(nums); maxIndex++ {
				if nums[maxIndex] != target {
					break
				}
			}
			maxIndex--

			return []int{minIndex, maxIndex}

		}
	}

	return []int{-1, -1}
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	res := searchRange([]int{5, 7, 7, 8, 8, 10}, 6)
	fmt.Print(res)
}
