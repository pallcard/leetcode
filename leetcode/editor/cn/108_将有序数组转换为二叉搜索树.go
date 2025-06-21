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
func sortedArrayToBST(nums []int) *TreeNode {

	var array2bst func(nums []int, left, right int) *TreeNode

	array2bst = func(nums []int, left, right int) *TreeNode {
		if left > right {
			return nil
		}
		mid := left + (right-left)/2
		node := &TreeNode{Val: nums[mid]}
		node.Left = array2bst(nums, left, mid-1)
		node.Right = array2bst(nums, mid+1, right)
		return node
	}
	return array2bst(nums, 0, len(nums)-1)
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
