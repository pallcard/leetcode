package main

import "fmt"

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}
	virtualHead := &ListNode{}

	ph := virtualHead
	pt := virtualHead
	q := head

	for q != nil {
		if q.Val >= x {
			pt.Next = q
			q = q.Next
			pt = pt.Next
			pt.Next = nil //断开原链表
		} else {
			phNext := ph.Next
			ph.Next = q
			q = q.Next
			ph = ph.Next
			ph.Next = phNext
			if phNext == nil {
				pt = ph
			}
		}

	}

	return virtualHead.Next

}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	arr := []int{1, 4, 3, 2, 5, 2}

	virtualHead := &ListNode{}

	p := virtualHead
	for _, v := range arr {
		p.Next = &ListNode{
			Val:  v,
			Next: nil,
		}
		p = p.Next
	}

	node := partition(virtualHead.Next, 3)
	fmt.Println(node)
}

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}
