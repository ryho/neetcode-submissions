/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type NodeWithI struct {
	node *TreeNode
	i int
}
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var stack []*NodeWithI

	var maxDepth int

	stack = append(stack, &NodeWithI{root,1})
	for len(stack) > 0 {
		tail := stack[len(stack)-1]
		stack = stack[0:len(stack)-1]
		if tail.i > maxDepth {
			maxDepth = tail.i
		}
		newI := tail.i+1
		if tail.node.Left != nil {
			stack = append(stack,&NodeWithI{tail.node.Left,newI})
		}
		if tail.node.Right != nil {
			stack = append(stack,&NodeWithI{tail.node.Right,newI})
		}
	}
	return maxDepth
}

func maxDepthRecursive(root *TreeNode) int {
    return maxDepthWithCounter(root,0)
}

func maxDepthWithCounter(root *TreeNode, i int) int {
	if root == nil {
		return i
	}
	return max(maxDepthWithCounter(root.Left,i+1),
	maxDepthWithCounter(root.Right,i+1))
}
