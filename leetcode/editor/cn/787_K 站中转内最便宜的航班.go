package main

import (
	"fmt"
	"math"
)

// leetcode submit region begin(Prohibit modification and deletion)
func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {

	flightMap := map[int][][]int{} //src:[](dst,w)
	for i := 0; i < len(flights); i++ {
		flightList, ok := flightMap[flights[i][0]]
		if !ok {
			flightList = make([][]int, 0)
		}
		flightList = append(flightList, []int{flights[i][1], flights[i][2]})
		flightMap[flights[i][0]] = flightList
	}

	memo := make([][]int, n)
	for i := 0; i < n; i++ {
		memo[i] = make([]int, k+1)
	}

	var dp func(cur, step int) int

	dp = func(cur, step int) int {
		if cur == dst {
			return 0
		}
		if step >= k+1 {
			return -1
		}

		if memo[cur][step] != 0 {
			return memo[cur][step]
		}

		res := math.MaxInt
		for _, flight := range flightMap[cur] {
			next := flight[0]
			w := flight[1]

			sub := dp(next, step+1)
			if sub == -1 {
				continue
			}
			if res > sub+w {
				res = sub + w
			}

		}
		if res == math.MaxInt {
			res = -1
		}
		memo[cur][step] = res
		return res

	}

	price := dp(src, 0)
	if price == math.MaxInt {
		return -1
	}

	return price
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	res := findCheapestPrice(4, [][]int{{0, 1, 100}, {1, 2, 100}, {2, 0, 100}, {1, 3, 600}, {2, 3, 200}}, 0, 3, 1)
	fmt.Print(res)
}
