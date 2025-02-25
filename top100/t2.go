package top100

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	p := l1
	q := l2

	res := new(ListNode)
	pre := res //前一个节点
	cur := res //当前节点

	for p != nil || q != nil {
		curV := cur.Val
		if p != nil {
			curV += p.Val
			p = p.Next
		}
		if q != nil {
			curV += q.Val
			q = q.Next
		}
		cur.Val = curV % 10
		nextNode := new(ListNode)
		nextNode.Val = curV / 10
		cur.Next = nextNode

		pre = cur
		// next
		cur = cur.Next
	}

	if cur.Val == 0 { //最后一个指针置空
		pre.Next = nil
	}

	return res

}
