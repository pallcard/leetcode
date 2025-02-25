package top100

// twoSum
func twoSum(nums []int, target int) []int {

	restMap := map[int]int{} //rest:index
	for i, num := range nums {
		if index, ok := restMap[num]; ok {
			return []int{index, i}
		} else {
			restMap[target-num] = i
		}
	}
	return []int{}
}
