/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    var currentHead *ListNode
	var newHead *ListNode
	for list1 != nil || list2 != nil {
		var next *ListNode
		if list1 == nil && list2 != nil {
			next = list2
			list2 = list2.Next
		} else if list1 != nil && list2 == nil {
			next = list1
			list1 = list1.Next
		} else if list1.Val < list2.Val {
			next = list1
			list1 = list1.Next
		} else {
			next = list2
			list2 = list2.Next
		}
		if currentHead != nil {
			currentHead.Next = next
		} else {
			newHead = next
		}
		currentHead = next
	}
	return newHead
}
