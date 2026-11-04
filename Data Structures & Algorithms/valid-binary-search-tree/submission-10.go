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
		// Left needs to be less than the min val we have seen to the right
		// The left most node in each row will still have MinInt here
		if !(root.Left.Val > minVal) {
			return false
		}
		if !isValidBSTHelper(root.Left, minVal,min(maxVal,root.Val)) {
			return false
		}
	}
    if root.Right != nil {
		// Right needs to be greater than root
		if !(root.Right.Val > root.Val) {
			return false
		}
		// Right needs to be greater than the max value we have seen to the left
		// for the farthest right, this will still be MaxInt
		if !(root.Right.Val < maxVal){
			return false
		}
		if !isValidBSTHelper(root.Right,max(minVal, root.Val),maxVal) {
			return false
		}
	}
	return true
}
