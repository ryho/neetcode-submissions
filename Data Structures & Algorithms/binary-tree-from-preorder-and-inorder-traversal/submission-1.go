/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var dataToInOrderIndex map[int]int
var inorderOriginal []int

func buildTree(preorder []int, inorder []int) *TreeNode {
	inorderOriginal = inorder
	dataToInOrderIndex = map[int]int{}
	for i,val := range inorder {
		dataToInOrderIndex[val] = i
	}
	preorder, head := buildTreeHelper(preorder, 0, len(inorder))
	if len(preorder) > 0 {
		panic("did not finish")
	}
	return head
}

func buildTreeHelper(preorder []int, x, y int) ([]int, *TreeNode) {
	val, preorder := preorder[0], preorder[1:]
    node := &TreeNode {
		Val: val,
	}
	i, j, k, l := split(val, x, y)
	if j-i > 0 {
		preorder, node.Left = buildTreeHelper(preorder, i, j)
	}
	if l-k > 0 {
		preorder, node.Right = buildTreeHelper(preorder, k, l)
	}
	
	return preorder, node
}

func split(num int, x, y int) (i,j,k,l int) {
	index, ok := dataToInOrderIndex[num]
	if !ok {
		panic("split: did not find num in map")
	}
	return x,index,index+1,y
}