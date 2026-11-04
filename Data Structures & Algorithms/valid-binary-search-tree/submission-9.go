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
	if root == nil {
		return true
	}
    if root.Left != nil {
		// Left needs to be less than root
		if !(root.Left.Val < root.Val) {
			return false
		}
		// left needs to be in the bounds
		if !(root.Left.Val > minVal) {
			return false
		}
		if !isValidBSTHelper(root.Left, minVal,min(maxVal,root.Val)) {
			return false
		}
	}
    if root.Right != nil {
		// right needs to be greater than root
		if !(root.Right.Val > root.Val) {
			return false
		}
		if !(root.Right.Val < maxVal){
			return false
		}
		if !isValidBSTHelper(root.Right,max(minVal, root.Val),maxVal) {
			return false
		}
	}
	return true
}
