package main

import (
	"fmt"
	"strconv"
)

// leetcode submit region begin(Prohibit modification and deletion)
func monotoneIncreasingDigits(n int) int {
	//todo  找到首个逆序的数-1，以此像前
	// 332 --> 329  --> 299

	nstr := strconv.Itoa(n)
	length := len(nstr)
	for i := length - 2; i >= 0; i-- {
		if nstr[i] > nstr[i+1] {
			//当前位置-1
			nstr = nstr[:i] + string(nstr[i]-1)
			// 后续位置变9
			str9 := ""
			for j := i + 1; j < length; j++ { //todo 注意使用length，不能使用len(nstr)，nstr大小改了
				str9 += "9"
			}
			nstr += str9
		}
	}

	res, _ := strconv.Atoi(nstr)
	return res
	//
	//nStr := fmt.Sprintf("%d", n)
	//pre := nStr[0]
	//index := 0
	//for i := 1; i < len(nStr); i++ {
	//	if nStr[i] < pre {
	//		index = i
	//		break
	//	}
	//}
	//if index == 0 {
	//	return n
	//}
	//str := ""
	//for i := index - 1; i >= 0; i-- {
	//	// nStr[i+1] = 9
	//	//str += "9"
	//	str = string(nStr[i]-1) + "9"
	//	//nStr = nStr[:i+1] + "9" + nStr[i+1:]
	//	// nStr[i] -= 1
	//	if i-2 >= 0 {
	//		nStr = nStr[:i-2] + string(nStr[i-1]-1) + nStr[i:]
	//	} else {
	//		nStr = string(nStr[i-1]-1) + nStr[i:]
	//	}
	//
	//}
	//
	//for i := n; i >= 0; i-- {
	//	if isIncr(i) {
	//		return i
	//	}
	//}
	//return 0
}

func isIncr(n int) bool {
	pre := 10
	for n > 0 {
		t := n % 10
		if t <= pre {
			pre = t
		} else {
			return false
		}
		n /= 10
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	digits := monotoneIncreasingDigits(10)
	fmt.Print(digits)
}
