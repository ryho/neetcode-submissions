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
	var arrays [][]int
	var currentArray []int
	rowLen := len(queue)
	i := 0
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		currentArray = append(currentArray, node.Val)
		if node.Left!= nil {
			queue = append(queue, node.Left)
		}
		if node.Right!= nil {
			queue = append(queue, node.Right)
		}
		i++
		if i == rowLen {
			arrays = append(arrays, currentArray)
			currentArray = nil
			i = 0
			rowLen = len(queue)
		}
	}
	return arrays
}

func levelOrderOld(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
    queue := []*TreeNode{root}
	var nextQueue []*TreeNode
	var arrays [][]int
	var currentArray []int
	for len(queue) > 0 || len(nextQueue) > 0 {
		if len(queue) == 0 {
			swap := queue
			queue = nextQueue
			nextQueue = swap
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
