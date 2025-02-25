package main
//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func kthToLast(head *ListNode, k int) int {
	if head == nil {
		return -1
	}
	p1 := head
	p2 := head
	for i := 1; i < k; i++ {
		if p1.Next == nil {
			return -1
		}else {
			p1 = p1.Next
		}
	}

	for ;p1.Next != nil; {
		p1 = p1.Next
		p2 = p2.Next
	}
	return p2.Val
}
//leetcode submit region end(Prohibit modification and deletion)


func main() {

}

