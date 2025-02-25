package main

// leetcode submit region begin(Prohibit modification and deletion)
func twoSum(nums []int, target int) []int {
	rest := map[int]int{}
	for i, num := range nums {
		if index,ok := rest[target- num]; ok {
			return []int{i, index}
		} else {
			rest[num]=i
		}
	}
	return []int{}
}

//leetcode submit region end(Prohibit modification and deletion)
