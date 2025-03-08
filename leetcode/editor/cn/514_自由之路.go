package main

import (
	"fmt"
	"math"
)

// leetcode submit region begin(Prohibit modification and deletion)
func findRotateSteps(ring string, key string) int {

	ringMap := map[byte][]int{}

	for i := 0; i < len(ring); i++ {
		indexArr, ok := ringMap[ring[i]]
		if !ok {
			indexArr = make([]int, 0)
		}
		indexArr = append(indexArr, i)
		ringMap[ring[i]] = indexArr
	}
	memo := make([][]int, len(ring))
	for i := 0; i < len(ring); i++ {
		memo[i] = make([]int, len(key))
	}

	var dp func(ringIndex, keyIndex int) int
	dp = func(ringIndex, keyIndex int) int {
		if keyIndex == len(key) {
			return 0
		}
		if memo[ringIndex][keyIndex] != 0 {
			return memo[ringIndex][keyIndex]
		}

		res := math.MaxInt
		for _, index := range ringMap[key[keyIndex]] {
			mStep := math.Abs(float64(ringIndex - index))
			mStep = math.Min(mStep, float64(len(ring))-mStep)
			sub := dp(index, keyIndex+1)
			if res > int(mStep)+sub+1 {
				res = int(mStep) + sub + 1
			}
		}
		memo[ringIndex][keyIndex] = res
		return res

	}

	return dp(0, 0)

}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	steps := findRotateSteps("godding", "gd")
	fmt.Print(steps)
}
