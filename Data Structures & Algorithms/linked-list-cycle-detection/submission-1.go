/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
	fast := head
	slow := head
	for fast != nil {
		next := fast.Next
		if next == nil {
			return false
		}
		fast = next.Next

		slow = slow.Next

		if fast == slow {
			return true
		}
	}
	return false
}
