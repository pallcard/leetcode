package main

// leetcode submit region begin(Prohibit modification and deletion)
func validMountainArray(arr []int) bool {
	if len(arr) <= 2 {
		return false
	}
	flag := 0 //
	pre := arr[0]
	for i := 1; i < len(arr); i++ {
		if pre < arr[i] {
			pre = arr[i]
			if flag == 0 { //前一次初始
				flag = 1
			} else if flag == 2 { //前一次下降
				return false
			}
		} else if pre > arr[i] {
			pre = arr[i]
			if flag == 1 { //前一次上升
				flag = 2
			} else if flag == 0 { //前一次初始
				return false
			}

		} else {
			return false
		}
	}

	return flag == 2

}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
