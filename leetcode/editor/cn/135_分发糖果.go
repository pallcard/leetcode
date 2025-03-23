package main

import "fmt"

// leetcode submit region begin(Prohibit modification and deletion)
func candy(ratings []int) int {
	//todo 1. 先从左到右给糖果
	// 2. 从又到左ckeck一遍

	candyList := make([]int, len(ratings), len(ratings))
	for i := 0; i < len(candyList); i++ {
		candyList[i] = 1
	}

	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			candyList[i] = candyList[i-1] + 1
		}
	}

	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] { //左边大时，判断饼干数是否大于右边
			if candyList[i] <= candyList[i+1] {
				candyList[i] = candyList[i+1] + 1
			}
		}
	}

	sum := 0
	for i := 0; i < len(candyList); i++ {
		sum += candyList[i]
	}
	return sum
}

func candy2(ratings []int) int {
	// todo *** 错误代码
	minIndex := 0
	minScore := ratings[0]
	for i := 1; i < len(ratings); i++ {
		if ratings[i] < minScore {
			minScore = ratings[i]
			minIndex = i
		}
	}

	candyCnt := 1
	pre := 1
	for i := minIndex + 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			candyCnt += pre + 1
			pre++
		} else {
			candyCnt += 1
			pre = 1
		}
	}

	pre = 1
	for i := minIndex - 1; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			candyCnt += pre + 1
			pre++
		} else {
			candyCnt += 1
			pre = 1
		}
	}

	return candyCnt

}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	res := candy([]int{1, 0, 2})
	fmt.Print(res)
}
