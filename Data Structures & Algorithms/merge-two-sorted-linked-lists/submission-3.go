/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	newHead := &ListNode{}
	current := newHead
	for list1 != nil && list2 != nil {
		var next *ListNode
		if list1.Val < list2.Val {
			next = list1
			list1 = list1.Next
		} else {
			next = list2
			list2 = list2.Next
		}
		current.Next = next
		current = next
	}
	current.Next = list1
	if list1 == nil {
		current.Next = list2
	}
	return newHead.Next
}
