/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
    _, ok := isBalancedRecursiveHelper(root)
	return ok
}

func isBalancedRecursiveHelper(node *TreeNode) (int, bool) {
	if node == nil {
		return 0, true
	}
	leftHeight, leftOk := isBalancedRecursiveHelper(node.Left)
	rightHeight, rightOk := isBalancedRecursiveHelper(node.Right)
	if !leftOk || !rightOk {
		return 0, false
	}
	diff := leftHeight-rightHeight
	if diff > 1 || diff < -1 {
		return 0, false
	}
	return max(leftHeight,rightHeight)+1, true
}