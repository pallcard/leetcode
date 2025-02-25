package top100

import (
	"fmt"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	l1 := genList([]int{2, 4, 9})
	l2 := genList([]int{5, 6, 4, 9})

	res := addTwoNumbers(l1, l2)
	fmt.Sprint(res)

}

func genList(arr []int) *ListNode {
	l1 := new(ListNode)
	curL1 := l1
	for i, v := range arr {
		if i > 0 {
			next := new(ListNode)
			next.Val = v
			curL1.Next = next
			curL1 = curL1.Next
		} else {
			curL1.Val = v
		}
	}
	return l1
}
