/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeKLists(lists []*ListNode) *ListNode {
    dummy := &ListNode{
		Val: math.MaxInt,
	}
	current := dummy
	for {
		minValue := math.MaxInt
		var minIndex int
		for i, list := range lists {
			if list != nil && list.Val < minValue {
				minValue = list.Val
				minIndex = i
				if minValue == current.Val {
					break
				}
			}
		}
		if minValue == math.MaxInt {
			break
		}
		current.Next = lists[minIndex]
		current = current.Next
		lists[minIndex] = lists[minIndex].Next
		if minIndex == 0 && lists[minIndex]==nil{
			lists = lists[1:]
		}
	}
	return dummy.Next
}
