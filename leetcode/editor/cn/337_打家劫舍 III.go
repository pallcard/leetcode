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

func rob(root *TreeNode) int {
	memo := map[*TreeNode]int{}
	return robWithMemo(root, memo)

}

func robWithMemo(root *TreeNode, memo map[*TreeNode]int) int {
	if root == nil {
		return 0
	}

	if memoRes, ok := memo[root]; ok {
		return memoRes
	}

	do := root.Val
	if root.Left != nil {
		do += robWithMemo(root.Left.Left, memo) + robWithMemo(root.Left.Right, memo)
	}
	if root.Right != nil {
		do += robWithMemo(root.Right.Left, memo) + robWithMemo(root.Right.Right, memo)
	}

	notdo := robWithMemo(root.Left, memo) + robWithMemo(root.Right, memo)

	res := max(do, notdo)
	memo[root] = res

	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
