package main

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	// todo 中序遍历
	var pre *TreeNode // 记录前一个节点

	var inOrder func(root *TreeNode) bool

	inOrder = func(root *TreeNode) bool {
		if root == nil {
			return true
		}
		if !inOrder(root.Left) {
			return false
		}

		if pre != nil && pre.Val >= root.Val {
			return false
		}
		pre = root //给前一个节点赋值

		if !inOrder(root.Right) {
			return false
		}
		return true

	}

	//if root == nil {
	//	return true
	//}
	//
	//var inOrder func(root *TreeNode) int
	//
	//isValid := true
	//inOrder = func(root *TreeNode) int {
	//	if root.Left != nil {
	//		left := inOrder(root.Left)
	//		if left >= root.Val {
	//			isValid = false
	//		}
	//	}
	//
	//	if root.Right != nil {
	//		right := inOrder(root.Right)
	//		if root.Val >= right {
	//			isValid = false
	//		}
	//	}
	//
	//	return root.Val
	//}

	return inOrder(root)
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
