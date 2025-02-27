package main

// leetcode submit region begin(Prohibit modification and deletion)
func twoSum(numbers []int, target int) []int {
	numMap := map[int]int{} // num:index

	for i, number := range numbers {
		if index, ok := numMap[target-number]; ok {
			return []int{index + 1, i + 1}
		} else {
			numMap[number] = i
		}
	}
	return []int{-1, -1}
}

func twoSum2(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1

	for left < right {
		if numbers[left]+numbers[right] == target {
			return []int{left + 1, right + 1}
		} else if numbers[left]+numbers[right] > target {
			right--
		} else {
			left++
		}
	}

	return []int{-1, -1}
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
