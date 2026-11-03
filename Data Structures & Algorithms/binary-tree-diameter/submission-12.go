/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var queue []*TreeNode
	data := map[*TreeNode][2]int{}
	data[nil] = [2]int{0,0}
	queue = append(queue, root)
	for len(queue) > 0 {
		node := queue[len(queue)-1]
		shouldContinue := false
		if node.Left != nil {
			_, ok := data[node.Left]
			if !ok {
				queue = append(queue, node.Left)
				shouldContinue = true
			}
		}
		if node.Right != nil {
			_, ok := data[node.Right]
			if !ok {
				queue = append(queue, node.Right)
				shouldContinue = true
			}
		}
		if shouldContinue {
			continue
		}
		queue = queue[:len(queue)-1]
		lData := data[node.Left]
		rData := data[node.Right]

		currentData := [2]int{max(lData[0], rData[0])+1, max(lData[1],rData[1],lData[0]+rData[0])}
		data[node] = currentData
	}
	result := data[root]
	return result[1]
}

func diameterOfBinaryTreeRecursive(root *TreeNode) int {
	_, maxSoFar := diameterOfBinaryTreeHelperRecursive(root)

	return maxSoFar
}

func diameterOfBinaryTreeHelperRecursive(root *TreeNode) (int,int) {
	if root == nil {
		return 0,0
	}
	left,maxSoFarL := diameterOfBinaryTreeHelperRecursive(root.Left)
	right,maxSoFarR := diameterOfBinaryTreeHelperRecursive(root.Right)

	return max(left, right)+1, max(maxSoFarL,maxSoFarR,left+right)
}
