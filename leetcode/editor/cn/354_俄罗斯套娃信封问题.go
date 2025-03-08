package main

import "sort"

// leetcode submit region begin(Prohibit modification and deletion)
func maxEnvelopes(envelopes [][]int) int {
	// todo 通过排序固定住一个纬度，处理成但纬度数组
	// todo 按照w升序，w相同的h降序 ===>  保证相同w的数据只取最大的数，从而排除掉相同w

	if len(envelopes) == 1 {
		return 1
	}

	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] < envelopes[j][0] { //w大的放后面，增序
			return true
		} else if envelopes[i][0] > envelopes[j][0] {
			return false
		} else if envelopes[i][1] > envelopes[j][1] { //h大的放前面，倒叙
			return true
		} else {
			return false
		}
	})

	nums := make([]int, 0, len(envelopes))
	for i := 0; i < len(envelopes); i++ {
		nums = append(nums, envelopes[i][1])
	}

	//dp := make([]int, len(envelopes), len(envelopes))
	//dp[0] = 1
	//for i := 1; i < len(dp); i++ {
	//	dp[i] = 1
	//	for j := 0; j < i; j++ {
	//		if envelopes[i][1] > envelopes[j][1] && dp[i] < dp[j]+1 {
	//			dp[i] = dp[j] + 1
	//		}
	//	}
	//}
	//max := 0
	//for i := 0; i < len(dp); i++ {
	//	if max < dp[i] {
	//		max = dp[i]
	//	}
	//}

	return lengthOfLIS3(nums)
}

func lengthOfLIS3(nums []int) int {
	// todo 按照蜘蛛纸牌的玩法吧数组分成n堆，每一堆都是从大到小的顺序，堆数则为最大的子序列
	// todo 每一堆都是由大到小，遇到无法放的则开一个新堆，则可以保证每一堆的最小的一个可以代表该堆了，
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

}
