package main

import "fmt"

// leetcode submit region begin(Prohibit modification and deletion)
func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	preIndex := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[preIndex] = nums[i]
			preIndex++
		}
	}

	return preIndex

}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

	element := removeElement([]int{3, 2, 2, 3}, 3)
	fmt.Println(element)
}
