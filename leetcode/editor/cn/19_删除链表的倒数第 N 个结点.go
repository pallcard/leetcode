package main

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}

	virtualHead := &ListNode{Next: head}

	p1 := virtualHead
	p2 := virtualHead
	for i := 0; i < n+1; i++ {
		p1 = p1.Next
	}

	for ; p1 != nil; {
		p1 = p1.Next
		p2 = p2.Next
	}

	p2.Next = p2.Next.Next
	return virtualHead.Next
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
