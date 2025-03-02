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
func preorderTraversal(root *TreeNode) []int {
	t := make([]int, 0)
	preTraver(root, &t)
	return t
}

func preTraver(root *TreeNode, track *[]int) {
	if root == nil {
		return
	}

	*track = append(*track, root.Val)
	if root.Left != nil {
		preTraver(root.Left, track)
	}

	if root.Right != nil {
		preTraver(root.Right, track)
	}

}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
