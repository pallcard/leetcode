package main
//leetcode submit region begin(Prohibit modification and deletion)
func twoSum(price []int, target int) []int {
	if len(price) <2 {
		return []int{-1,-1}
	}

	left := 0
	right := len(price)-1

	for left < right {
		if price[left] + price[right] == target {
			return []int{ price[left], price[right]}
		}  else if  price[left] + price[right] > target {
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

