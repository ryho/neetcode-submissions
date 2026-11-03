/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	return isSubtreeHelperNotStarted(root,subRoot)
}

func isSubtreeHelperStarted(root, subRoot *TreeNode) bool {
	if root == nil && subRoot == nil {
		return true
	}
	if root == nil || subRoot == nil {
		return false
	}
	if root.Val != subRoot.Val {
		return false 
	}
	left := isSubtreeHelperStarted(root.Left, subRoot.Left)
	right := isSubtreeHelperStarted(root.Right, subRoot.Right)
	return left && right
}

func isSubtreeHelperNotStarted(root, originalSubRoot *TreeNode) bool {
	if root == nil && originalSubRoot == nil {
		return true
	}
	if root == nil || originalSubRoot == nil {
		return false
	}
	if root.Val == originalSubRoot.Val {
		left := isSubtreeHelperStarted(root.Left, originalSubRoot.Left) 
		right := isSubtreeHelperStarted(root.Right, originalSubRoot.Right)
		if left && right {
			return true
		}
	}
	left := isSubtreeHelperNotStarted(root.Left, originalSubRoot)
	right := isSubtreeHelperNotStarted(root.Right, originalSubRoot)
	return left || right
}
