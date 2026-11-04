/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func goodNodes(root *TreeNode) int {
    return goodNodesHelper(root, math.MinInt)
}

func goodNodesHelper(node *TreeNode, maxVal int) int {
	if node == nil {
		return 0
	}
	maxVal = max(node.Val, maxVal)
	leftCount := goodNodesHelper(node.Left, maxVal)
	rightCount := goodNodesHelper(node.Right, maxVal)
	count := leftCount + rightCount
	if node.Val >= maxVal {
		count++
	}
	return count
}