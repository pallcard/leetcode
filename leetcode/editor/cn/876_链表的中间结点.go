package main

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
	fast := head
	slow := head

	for fast != nil {
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		} else {
			return slow
		}
		slow = slow.Next
	}

	return slow
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
