package main

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	slow := head
	fast := head

	isCycle := false
	for fast != nil && fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if slow == fast {
			isCycle = true
			break
		}
	}

	if !isCycle {
		return nil
	}

	if slow == head {
		return slow
	}
	p := head
	for {
		p = p.Next
		slow = slow.Next
		if p == slow {
			return p
		}
	}

}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
