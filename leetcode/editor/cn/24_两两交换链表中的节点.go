package main

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	newHead := &ListNode{}
	p := newHead

	q := head
	for q != nil {
		q1 := q.Next
		if q1 != nil {
			var r *ListNode
			if q1.Next != nil {
				r = q1.Next
			}

			p.Next = q1
			q1.Next = q
			p = q

			q = r
		} else {
			break
		}

	}

	p.Next = q
	return newHead.Next
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
