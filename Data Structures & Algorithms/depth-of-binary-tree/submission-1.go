/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxDepth(root *TreeNode) int {
    return maxDepthWithCounter(root,0)
}

func maxDepthWithCounter(root *TreeNode, i int) int {
	if root == nil {
		return i
	}
	return max(maxDepthWithCounter(root.Left,i+1),
	maxDepthWithCounter(root.Right,i+1))
}
