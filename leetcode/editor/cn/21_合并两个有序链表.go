package main

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    virtualHead := &ListNode{}
	r := virtualHead
	p := list1
	q := list2
	for ;p != nil && q != nil; {
		if p.Val <= q.Val {
			r.Next = p
			p = p.Next
		} else {
			r.Next = q
			q = q.Next
		}
		r = r.Next
	}
	for ;p != nil; {
		r.Next = p
		p = p.Next
		r = r.Next
	}
	for ;q != nil; {
		r.Next = q
		q = q.Next
		r = r.Next
	}
	return virtualHead.Next

}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
