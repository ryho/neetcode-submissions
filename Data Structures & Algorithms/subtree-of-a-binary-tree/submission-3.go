/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	return searchForSameTree(root,subRoot)
}

func sameTree(root, subRoot *TreeNode) bool {
	if root == nil && subRoot == nil {
		return true
	}
	if root == nil || subRoot == nil {
		return false
	}
	if root.Val != subRoot.Val {
		return false 
	}
	left := sameTree(root.Left, subRoot.Left)
	right := sameTree(root.Right, subRoot.Right)
	return left && right
}

func searchForSameTree(root, originalSubRoot *TreeNode) bool {
	if root == nil && originalSubRoot == nil {
		return true
	}
	if root == nil || originalSubRoot == nil {
		return false
	}
	if root.Val == originalSubRoot.Val {
		left := sameTree(root.Left, originalSubRoot.Left) 
		right := sameTree(root.Right, originalSubRoot.Right)
		if left && right {
			return true
		}
	}
	left := searchForSameTree(root.Left, originalSubRoot)
	right := searchForSameTree(root.Right, originalSubRoot)
	return left || right
}
