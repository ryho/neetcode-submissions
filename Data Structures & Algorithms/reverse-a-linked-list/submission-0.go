/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
	var previous *ListNode
	var lastHead *ListNode
	for head != nil {
		lastHead = head
		second := head.Next
		var third *ListNode
		if second != nil {
			third = second.Next
		}

		head.Next = previous
		if second != nil {
			second.Next = head
			lastHead = second
		}

		head = third
		previous = second
	}
	return lastHead
}
