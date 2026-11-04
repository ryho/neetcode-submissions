/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var maxSoFar int
func maxPathSum(root *TreeNode) int {
    maxSoFar = math.MinInt
	maxPathSumHelper(root)
	return maxSoFar
}

func maxPathSumHelper(node *TreeNode) int {
    if node == nil {
		return math.MinInt32
	}
	left := maxPathSumHelper(node.Left)
	right := maxPathSumHelper(node.Right)
	value := max(left+node.Val, right+node.Val, node.Val)

	maxSoFar = max(maxSoFar, value, left+node.Val+right, left, right)
	return value
}
