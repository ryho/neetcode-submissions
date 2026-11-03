/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	var stack []*TreeNode
	stack = append(stack, root)
	for len(stack)>0 {
		tail := stack[len(stack)-1]
		temp := tail.Left
		tail.Left = tail.Right
		tail.Right = temp
		stack = stack[0:len(stack)-1]
		if tail.Left != nil {
			stack = append(stack, tail.Left)
		}
		if tail.Right != nil {
			stack = append(stack, tail.Right)
		}
	}
	return root
}

func invertTreeRecursiveDFS(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	invertTree(root.Left)
	invertTree(root.Right)
	temp := root.Left
	root.Left = root.Right
	root.Right = temp
	return root
}
