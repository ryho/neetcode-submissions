/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Entry struct {
	node *TreeNode
	maxVal int
}
func goodNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var goodCount int
	queue :=[]*Entry{
		&Entry{
			node:root,
			maxVal:math.MinInt,
		},
	}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		maxVal := max(node.node.Val, node.maxVal)
		if node.node.Val >= maxVal {
			goodCount++
		}
		if node.node.Left != nil {
			queue = append(queue,
				&Entry{
					node:node.node.Left,
					maxVal:maxVal,
				})
		}
		if node.node.Right != nil {
			queue = append(queue,
				&Entry{
					node:node.node.Right,
					maxVal: maxVal,
				})
		}
	}

	return goodCount
}

func goodNodesRecursive(root *TreeNode) int {
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