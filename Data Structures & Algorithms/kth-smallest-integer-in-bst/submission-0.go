/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthSmallest(root *TreeNode, k int) int {
    value, _ := kthSmallestHelper(root, k,nil)
	if value == nil {
		panic("did not find")
	}
	return *value
}

func kthSmallestHelper(node *TreeNode, k int, count *int) (*int, *int) {
	if node.Left != nil {
		var value *int
		value, count = kthSmallestHelper(node.Left, k, count)
		if value != nil {
			return value, nil
		}
	} else if count == nil {
		count = ptr(0)
	}
	*count++
	if *count == k {
		return ptr(node.Val), nil
	}
	if node.Right != nil {
		value, count := kthSmallestHelper(node.Right, k, count)
		return value,count
	}
	return nil, count

}

func ptr(n int)*int{
	return &n
}

func dePtr(n *int) int {
	if n == nil {
		return math.MinInt
	}
	return *n
}

