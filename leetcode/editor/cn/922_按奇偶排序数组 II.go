package main

// leetcode submit region begin(Prohibit modification and deletion)
func sortArrayByParityII(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	i, j := 0, 1 //0偶数 1奇数
	for ; i < len(nums); i += 2 {
		if nums[i]%2 == 0 {
			continue
		} else {
			for ; j < len(nums); j += 2 {
				if nums[j]%2 != 1 {
					break
				}
			}

			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	return nums

}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
