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
func balanceBST(root *TreeNode) *TreeNode {
	//todo 1. 转化为增序数组 先序遍历
	//     2. 数组转平衡数 t108
	if root == nil {
		return root
	}
	arr := make([]int, 0)
	var preOder func(root *TreeNode)
	preOder = func(root *TreeNode) {
		if root == nil {
			return
		}
		preOder(root.Left)
		arr = append(arr, root.Val)
		preOder(root.Right)
	}
	preOder(root)

	var arr2bst func(arr []int, left, right int) *TreeNode
	arr2bst = func(arr []int, left, right int) *TreeNode {
		if left > right {
			return nil
		}

		mid := left + (right-left)/2
		node := &TreeNode{
			Val: arr[mid],
		}

		node.Left = arr2bst(arr, left, mid-1)
		node.Right = arr2bst(arr, mid+1, right)
		return node
	}
	return arr2bst(arr, 0, len(arr)-1)
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
