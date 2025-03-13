package main

import (
	"fmt"
	"strings"
)

// leetcode submit region begin(Prohibit modification and deletion)
func multiply(num1 string, num2 string) string {
	// 123
	// 456
	//todo 关键：num1[i1]*num2[i2]的结果正好在下标的 i1+i2, i1+i2+1
	/**
			1  2  3
		       4  5
		-------------
		       1  5
		    1  0
		 0  5
		    1  2
		 0  8
	  0	 4
	------------------
	  0  5   5  3   5
	i 0  1   2  3   4

		**/

	if num1 == "0" || num2 == "0" {
		return "0"
	}

	m := len(num1)
	n := len(num2)
	res := make([]int, m+n)

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			n1 := int(num1[i] - '0')
			n2 := int(num2[j] - '0')
			sum := n1*n2 + res[i+j+1]
			res[i+j+1] = sum % 10 // todo 地位 上一行以计算了低位
			res[i+j] += sum / 10  // todo 高位 +=

		}
	}

	numStr := strings.Builder{}
	for i := 0; i < len(res); i++ {
		numStr.WriteRune(rune(res[i] + '0'))
	}
	return strings.TrimLeft(numStr.String(), "0") //去除前缀0
}
func multiply2(num1 string, num2 string) string {

	if num1 == "0" || num2 == "0" {
		return "0"
	}

	//res := make([]byte, 0)
	res := ""
	for i := len(num1) - 1; i >= 0; i-- {
		n1 := num1[i] - '0'

		for j := len(num2) - 1; j >= 0; j-- {
			n2 := num2[j] - '0'
			temp := fmt.Sprintf("%d", n1*n2)

			for k := 0; k < len(num1)-1-i+len(num2)-1-j; k++ {
				temp += "0"
			}
			res = add(res, temp)

		}

	}

	return res

}

func add(num1 string, num2 string) string {
	i1 := len(num1) - 1
	i2 := len(num2) - 1

	res := make([]byte, 0)
	var h uint8 = 0
	var l uint8 = 0
	for i1 >= 0 || i2 >= 0 {
		var n1 uint8 = 0
		var n2 uint8 = 0
		if i1 >= 0 {
			n1 = num1[i1] - '0'
		}
		if i2 >= 0 {
			n2 = num2[i2] - '0'
		}
		l = (n1 + n2 + h) % 10
		h = (n1 + n2 + h) / 10
		res = append(res, l+'0')
		i1--
		i2--
	}

	if h > 0 {
		res = append(res, '1')
	}

	left := 0
	right := len(res) - 1
	for left < right {
		res[left], res[right] = res[right], res[left]
		left++
		right--
	}
	return string(res)
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	s := add("123", "456")
	fmt.Print(s)
	multiply("123", "45")
}
