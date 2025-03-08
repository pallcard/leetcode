package main
//leetcode submit region begin(Prohibit modification and deletion)
func subsets(nums []int) [][]int {
    res := make([][]int, 0)

	var backtrack func(trace []int, size,cur int)

	backtrack = func(trace []int, size,cur int) {
		//if size == len(trace) {
			tempArr := make([]int, 0, size)
			tempArr = append(tempArr, trace...)
			res = append(res, tempArr)
			//return
		//}

		for i := cur; i < len(nums); i++ {

			trace = append(trace, nums[i])
			backtrack(trace, size,i+1)
			trace = trace[:len(trace)-1]
		}
	}

	//for i := 0; i <= len(nums); i++ {
	//	backtrack([]int{}, i, 0)
		backtrack([]int{}, 0, 0)
	//}
	return res
}
//leetcode submit region end(Prohibit modification and deletion)


func main() {
	
}

