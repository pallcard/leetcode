package main

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func isPalindrome(head *ListNode) bool {
	// todo 1.使用快慢指针找到中点
	// todo 2.反转终点之后链表
	// todo 3.比较前半段 和 反转之后的后半段

	// <=1个节点
	if head == nil || head.Next == nil {
		return true
	}

	slow := head
	fast := head
	pre := new(ListNode) //记录pre为了分割链表，断开前后两段
	for fast != nil && fast.Next != nil {
		pre = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	pre.Next = nil

	p1 := head
	p2 := reverseList(slow)
	for p1 != nil && p1.Val == p2.Val {
		p1 = p1.Next
		p2 = p2.Next
	}

	if p1 == nil {
		return true
	}
	return false
}

func reverseList(head *ListNode) *ListNode {
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

func isPalindrome2(head *ListNode) bool {
	// todo 链表数据存入数组，判断数组是否回文
	arr := make([]int, 0)
	p := head
	for p != nil {
		arr = append(arr, p.Val)
		p = p.Next
	}

	left := 0
	right := len(arr) - 1

	for left < right && arr[left] == arr[right] {
		left++
		right--
	}

	if left >= right {
		return true
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1, Next: &ListNode{
				Val: 2, Next: &ListNode{
					Val: 1,
				},
			},
		},
	}
	isPalindrome(head)
}

type ListNode struct {
	Val  int
	Next *ListNode
}
