/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	queue := []*TreeNode{root}
	var result []int
	rowCount := len(queue)
	var i int

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}

		i++
		if i == rowCount {
			result = append(result, node.Val)
			i = 0 
			rowCount = len(queue)
		}
	}

	return result
}

func rightSideViewSilly(root *TreeNode) []int {
	var values []int
    for root != nil {
		values = append(values, root.Val)
		root = root.Right
	}
	return values
}
