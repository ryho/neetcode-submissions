/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	queue := []*TreeNode{root}
	data := map[*TreeNode]int{}
	data[nil] = 0
	for len(queue) > 0 {
		node := queue[len(queue)-1]
		var shouldContinue bool
		if _, ok := data[node.Left];!ok {
			queue = append(queue, node.Left)
			shouldContinue = true
		}
		if _, ok := data[node.Right];!ok {
			queue = append(queue, node.Right)
			shouldContinue = true
		}
		if shouldContinue {
			continue
		}
		queue = queue[:len(queue)-1]

		leftHeight := data[node.Left]
		rightHeight := data[node.Right]

		diff := leftHeight-rightHeight

		if diff > 1 || diff < -1 {
			return false
		}

		data[node] = max(leftHeight,rightHeight)+1

	}
	return true
}

func isBalancedRecursive(root *TreeNode) bool {
    _, ok := isBalancedRecursiveHelper(root)
	return ok
}

func isBalancedRecursiveHelper(node *TreeNode) (int, bool) {
	if node == nil {
		return 0, true
	}
	leftHeight, leftOk := isBalancedRecursiveHelper(node.Left)
	rightHeight, rightOk := isBalancedRecursiveHelper(node.Right)
	if !leftOk || !rightOk {
		return 0, false
	}
	diff := leftHeight-rightHeight
	if diff > 1 || diff < -1 {
		return 0, false
	}
	return max(leftHeight,rightHeight)+1, true
}