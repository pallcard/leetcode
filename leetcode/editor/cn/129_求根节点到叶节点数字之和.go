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
func sumNumbers(root *TreeNode) int {

	var dfs func(root *TreeNode, val int)
	sum := 0
	dfs = func(root *TreeNode, val int) {
		if root == nil {
			return
		}

		val = val*10 + root.Val
		if root.Left == nil && root.Right == nil {
			sum += val
			return
		}

		dfs(root.Left, val)
		dfs(root.Right, val)
	}

	dfs(root, 0)

	return sum
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
