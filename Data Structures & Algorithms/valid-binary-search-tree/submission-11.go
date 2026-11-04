/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isValidBSTHelper(root, math.MinInt, math.MaxInt)
}
func isValidBSTHelper(root *TreeNode, minVal, maxVal int) bool {
	if root.Val >= maxVal || root.Val <= minVal {
		return false
	}
    if root.Left != nil {
		if !isValidBSTHelper(root.Left, minVal, root.Val) {
			return false
		}
	}
    if root.Right != nil {
		if !isValidBSTHelper(root.Right, root.Val, maxVal) {
			return false
		}
	}
	return true
}
