/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 var maxSoFar int

func diameterOfBinaryTree(root *TreeNode) int {
	maxSoFar = 0
	diameterOfBinaryTreeHelper(root)

	fmt.Printf("maxSoFar %v\n", maxSoFar)

	return maxSoFar
}
func diameterOfBinaryTreeHelper(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := diameterOfBinaryTreeHelper(root.Left)
	right := diameterOfBinaryTreeHelper(root.Right)
	fmt.Printf("val %v left %v right %v\n", root.Val, left, right)
	if left + right > maxSoFar {
	fmt.Printf("maxSoFar b %v\n", maxSoFar)
		maxSoFar = left + right
	fmt.Printf("maxSoFar a %v\n", maxSoFar)
	}
	return max(left, right)+1


}
