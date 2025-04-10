package main

import "sort"

// leetcode submit region begin(Prohibit modification and deletion)
func smallerNumbersThanCurrent(nums []int) []int {

	// todo copy一份数组出来，排序， map[int]int value:排序后value的第一次出现的下标(正好也是该值对应的比他小的数的个数)

	res := make([]int, len(nums), len(nums))
	copy(res, nums)
	sort.Ints(res)
	resMap := map[int]int{}
	for i := 0; i < len(res); i++ {
		if _, ok := resMap[res[i]]; !ok {
			resMap[res[i]] = i
		}
	}

	for i := 0; i < len(nums); i++ {
		res[i] = resMap[nums[i]]
	}
	return res

	//res := make([]int, 0, len(nums))
	//for i := 0; i < len(nums); i++ {
	//	cnt := 0
	//	for j := 0; j < len(nums); j++ {
	//		if i != j && nums[j] < nums[i] {
	//			cnt++
	//		}
	//	}
	//	res = append(res, cnt)
	//}
	//
	//return res
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
