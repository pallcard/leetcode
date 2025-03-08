package main

import (
	"fmt"
)

// leetcode submit region begin(Prohibit modification and deletion)
func lengthOfLIS2(nums []int) int {
	// todo dp[i] : 以i结尾的最长的递增子序列

	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums), len(nums))
	dp[0] = 1

	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		for j := i - 1; j >= 0; j-- {
			if nums[i] > nums[j] && dp[i] < dp[j]+1 {
				dp[i] = dp[j] + 1
			}
		}
	}

	max := 0
	for i := 0; i < len(dp); i++ {
		if max < dp[i] {
			max = dp[i]
		}
	}
	return max

}

func lengthOfLIS(nums []int) int {
	// 按照蜘蛛纸牌的玩法吧数组分成n堆，每一堆都是从大到小的顺序，堆数则为最大的子序列
	// 每一堆都是由大到小，遇到无法放的则开一个新堆，则可以保证每一堆的最小的一个可以代表该堆了，
	top := make([]int, 0)

	for i := 0; i < len(nums); i++ {

		left := 0
		right := len(top)

		for left < right {
			mid := left + (right-left)/2
			if top[mid] >= nums[i] {
				right = mid
			} else if top[mid] < nums[i] {
				left = mid + 1
			}
		}

		if left == len(top) {
			top = append(top, nums[i])
		} else {
			top[left] = nums[i]
		}

	}

	return len(top)

}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	res := lengthOfLIS2([]int{1, 3, 6, 7, 9, 4, 10, 5, 6})
	fmt.Print(res)
}
