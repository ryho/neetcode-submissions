/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
    queue := []*TreeNode{root}
	var nextQueue []*TreeNode
	var arrays [][]int
	var currentArray []int
	for len(queue) > 0 || len(nextQueue) > 0{
		if len(queue) == 0 {
			queue = nextQueue
			nextQueue = nil
			arrays = append(arrays,currentArray)
			currentArray = nil
		}
		node := queue[0]
		queue = queue[1:]

		currentArray = append(currentArray, node.Val)
		if node.Left!= nil {
			nextQueue = append(nextQueue, node.Left)
		}
		if node.Right!= nil {
			nextQueue = append(nextQueue, node.Right)
		}
	}
	arrays = append(arrays, currentArray)
	return arrays
}
