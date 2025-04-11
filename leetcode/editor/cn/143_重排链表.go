package main

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	//todo 1.找中间节点
	//todo 2.反转链表
	//todo 3.交替拼链表

	slow := head
	fast := head

	var pre *ListNode
	for fast != nil && fast.Next != nil {
		pre = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	pre.Next = nil

	p1 := head
	p2 := reverseList2(slow)

	newHead := &ListNode{}
	p := newHead
	for p1 != nil && p2 != nil {
		p.Next = p1
		p1 = p1.Next
		p = p.Next

		p.Next = p2
		p2 = p2.Next
		p = p.Next
	}

	if p2 != nil {
		p.Next = p2
	}

}

func reverseList2(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {

		temp := cur
		cur = cur.Next

		temp.Next = pre
		pre = temp
	}
	return pre
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
