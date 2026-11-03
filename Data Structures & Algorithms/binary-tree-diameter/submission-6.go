/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


func diameterOfBinaryTree(root *TreeNode) int {
	_, maxSoFar := diameterOfBinaryTreeHelper(root)

	return maxSoFar
}
func diameterOfBinaryTreeHelper(root *TreeNode) (int,int) {
	if root == nil {
		return 0,0
	}
	left,maxSoFarL := diameterOfBinaryTreeHelper(root.Left)
	right,maxSoFarR := diameterOfBinaryTreeHelper(root.Right)

	return max(left, right)+1, max(maxSoFarL,maxSoFarR,left+right)


}
