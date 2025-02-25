package main
//leetcode submit region begin(Prohibit modification and deletion)
func containsDuplicate(nums []int) bool {
	repeat := map[int]struct{}{}
	for _, num := range nums {
		if _,ok := repeat[num]; ok {
			return true
		} else {
			repeat[num] = struct{}{}
		}
	}
	return false
}
//leetcode submit region end(Prohibit modification and deletion)
