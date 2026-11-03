/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	if p.Val > q.Val {
		swap := q
		q = p
		p = swap
	}
	return lowestCommonAncestorHelper(root, p, q)
}

func lowestCommonAncestorHelper(node *TreeNode, p *TreeNode, q *TreeNode) (*TreeNode) {
	if q.Val < node.Val {
		return lowestCommonAncestorHelper(node.Left,p,q)
	} else if p.Val>node.Val {
		return lowestCommonAncestorHelper(node.Right,p,q)
	}
	return node
}

