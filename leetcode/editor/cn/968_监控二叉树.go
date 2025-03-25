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
func minCameraCover(root *TreeNode) int {
	//todo 设置三种状态
	// 0. 无覆盖
	// 1. 有摄像头
	// 2. 有覆盖
	// 几种情况
	//（1）空节点,则有覆盖 ===> 2
	//（2）左右子节点均有覆盖,则该节点无覆盖，等待父节点来覆盖 ===> 0
	//（3）左右子节点有一个无覆盖，则该节点需要放摄像头 ====> 1
	//（3）左右子节点有一个有摄像头， ====> 2
	//（4）根节点无无覆盖，则需要放一个摄像头 ===> 1

	cnt := 0

	var trace func(root *TreeNode) int

	trace = func(root *TreeNode) int {
		if root == nil {
			return 2
		}
		left := trace(root.Left)
		right := trace(root.Right)

		if left == 2 && right == 2 {
			return 0
		}

		if left == 0 || right == 0 {
			cnt++
			return 1
		}

		if left == 1 || right == 1 {
			return 2
		}
		return -1
	}

	if trace(root) == 0 {
		cnt++
	}

	return cnt
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
