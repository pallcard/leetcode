package main
//leetcode submit region begin(Prohibit modification and deletion)
func twoSum(numbers []int, target int) []int {
	if len(numbers) <2 {
		return []int{-1,-1}
	}

	left := 0
	right := len(numbers)-1

	for left < right {
		if numbers[left] + numbers[right] == target {
			return []int{left, right}
		}  else if  numbers[left] + numbers[right] > target {
			right--
		} else {
			left++
		}
	}

	return  []int{-1,-1}
}
//leetcode submit region end(Prohibit modification and deletion)


func main() {
	
}

