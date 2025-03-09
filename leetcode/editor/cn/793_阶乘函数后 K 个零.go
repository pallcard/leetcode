package main

// leetcode submit region begin(Prohibit modification and deletion)
func preimageSizeFZF(k int) int {
	if k < 0 {
		return 0
	}
	if k == 0 {
		return 5
	}
	// todo 能够满足某个值x（x为非负整数）的阶乘，可以出现k个5的数量。
	//那么我们如果可以获得第一次满足5出现了k次，和第一次满足5出现k+1次，并且两者相减
	//结果就是满足了5可以出现k次的数量了。
	// 这个答案只有两种情况 0 或者 5, 出现k次时，到k+1次会间隔5个数字

	fzf := func(n int) int {
		cnt0 := 0
		for n > 0 {
			n /= 5
			cnt0 += n
		}
		return cnt0
	}

	// [0，5k]
	start := 0
	end := 5 * k
	for start <= end {
		mid := start + (end-start)/2

		cnt0 := fzf(mid)
		if cnt0 == k {
			return 5
		} else if cnt0 < k {
			start = mid + 1
		} else if cnt0 > k {
			end = mid - 1
		}
	}

	return 0

}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	preimageSizeFZF(3)
}
