/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func buildTree(preorder []int, inorder []int) *TreeNode {
	preorder, head := buildTreeHelper(preorder, inorder)
	if len(preorder) > 0 {
		panic("did not finish")
	}
	return head
}

func buildTreeHelper(preorder []int, inorder []int) ([]int, *TreeNode) {
	val, preorder := preorder[0], preorder[1:]
    node := &TreeNode {
		Val: val,
	}
	leftNums, rightNums := split(val, inorder)
	if len(leftNums) > 0 {
		preorder, node.Left = buildTreeHelper(preorder, leftNums)
	}
	if len(rightNums) > 0 {
		preorder, node.Right = buildTreeHelper(preorder, rightNums)
	}
	
	return preorder, node
}

func split(num int, nums []int) ([]int, []int) {
	for i :=0;i< len(nums); i++ {
		if nums[i]== num {
			return nums[:i],nums[i+1:]
		}
	}
	panic("split: did not find num in nums")
}