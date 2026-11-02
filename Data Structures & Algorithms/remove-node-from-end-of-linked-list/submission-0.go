/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    if head == nil {
		return nil
	}
	length := linkedListLength(head)
	if n > length {
		return head
	}
	indexToRemove := length-n

	if indexToRemove == 0 {
		return head.Next
	}
	cur := head
	for i := 0; i < indexToRemove-1; i++ {
		cur=cur.Next
	}
	cur.Next=cur.Next.Next
	return head
}

func linkedListLength(head *ListNode) int {
	var i int
	for head != nil {
		i++
		head = head.Next
	}
	return i
}
