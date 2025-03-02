package main

import "fmt"

// leetcode submit region begin(Prohibit modification and deletion)
func permute(nums []int) [][]int {
	res := make([][]int, 0)
	if len(nums) == 0 {
		return res
	}

	track := make([]int, 0)
	used := make([]bool, len(nums))
	backtrack(nums, track, used, &res) //dfs

	return res

}

func backtrack(nums []int, track []int, used []bool, res *[][]int) {
	if len(nums) == len(track) {
		copyTrack := make([]int, 0, len(track))
		copyTrack = append(copyTrack, track...)
		*res = append(*res, copyTrack) //res需要使用指针，否则地址值一直在被改变（输出的时候需要输出原切片）
	}
	for i, num := range nums {
		if used[i] {
			continue
		}
		track = append(track, num) //trace可以不用指针，trace的地址值一直改变不影响（输出的时候不需要输出原切片）
		used[i] = true
		backtrack(nums, track, used, res)
		used[i] = false
		track = track[:len(track)-1]
	}

}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	a := make([]int, 0)
	fmt.Printf("%p\n", a)
	a = append(a, 1)
	fmt.Printf("%p\n", a)
}
