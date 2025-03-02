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

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxDiameter := 0

	var maxDept func(root *TreeNode) int
	maxDept = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		left := maxDept(root.Left)
		right := maxDept(root.Right)

		if left+right > maxDiameter {
			maxDiameter = right + left
		}

		if left > right {
			return left + 1
		} else {
			return right + 1
		}
	}

	maxDept(root)
	return maxDiameter
}

func diameterOfBinaryTree2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxDiameter := 0

	var traverse func(root *TreeNode)
	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}

		left := maxDepth2(root.Left)   //计算节点深度
		right := maxDepth2(root.Right) //计算节点深度

		if (left + right) > maxDiameter {
			maxDiameter = left + right
		}

		traverse(root.Left) //遍历
		traverse(root.Right)
	}

	traverse(root)
	return maxDiameter
}

func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := maxDepth2(root.Left)
	rightDepth := maxDepth2(root.Right)

	if rightDepth > leftDepth {
		return rightDepth + 1
	}

	return leftDepth + 1
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
