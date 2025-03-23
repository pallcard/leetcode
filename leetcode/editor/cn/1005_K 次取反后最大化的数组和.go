package main

import (
	"fmt"
	"math"
	"sort"
)

// leetcode submit region begin(Prohibit modification and deletion)
func largestSumAfterKNegations(nums []int, k int) int {
	nega := make([]int, 0)
	sum := 0
	minPositive := math.MaxInt
	// todo 先加上所有正数
	for i := 0; i < len(nums); i++ {
		if nums[i] >= 0 {
			sum += nums[i]
			if nums[i] < minPositive {
				minPositive = nums[i]
			}
		} else {
			nega = append(nega, nums[i])
		}
	}
	sort.Ints(nega)
	i := 0
	// 处理负数
	for k > 0 && i < len(nega) {
		temp := -nega[i]
		if temp < minPositive {
			minPositive = temp
		}
		sum += temp
		k--
		i++
	}

	// todo *** 还存在负数需要加上
	for i < len(nega) {
		sum += nega[i]
		i++
	}

	if k%2 == 0 { //todo 包含k==0的情况
		return sum
	} else {
		return sum - minPositive*2 //todo ****需要*2
	}

}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	res := largestSumAfterKNegations([]int{4, 2, 3}, 1)
	fmt.Print(res)
}
