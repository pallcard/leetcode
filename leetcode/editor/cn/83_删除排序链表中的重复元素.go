package main
//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil {
		return nil
	}

	left := head
    right := head.Next

	for right != nil {
		if left.Val == right.Val {
			left.Next = right.Next
			right = left.Next
		} else {
			left = right
			right = right.Next

		}

	}
	return head

}
//leetcode submit region end(Prohibit modification and deletion)


func main() {
	
}

