package main

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	p1 := headA
	p2 := headB
	for p1 != p2 {
		if p1 != nil {
			p1 = p1.Next
		} else {
			p1 = headB
		}

		if p2 != nil {
			p2 = p2.Next
		} else {
			p2 = headA
		}

	}

	return p1
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
