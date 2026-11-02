/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy
	var carry int
	for l1 != nil || l2 != nil || carry > 0 {
		// fmt.Printf("l1 %v l2 %v carry %v current %v", l1, l2, carry, current)
		result := carry
		if l1 != nil {
			result += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			result += l2.Val
			l2 = l2.Next
		}
		
		currentDigit := result % 10
		carry = result / 10 

		newNode := &ListNode {
			Val: currentDigit,
		}
		current.Next = newNode
		current = newNode
	}

	return dummy.Next
}

func addTwoNumbersOld(l1 *ListNode, l2 *ListNode) *ListNode {
    l1Val := NodeToNumber(l1)
	l2Val := NodeToNumber(l2)

	result := l1Val + l2Val
	fmt.Printf("l1Val %v + l2Val %v = result %v", l1Val, l2Val, result)

	return EncodeNumber(result)
}


func NodeToNumber(l *ListNode) int64 {
	val := int64(0)
	cur := l
	mult := int64(1)
	for cur != nil {
		val = val+int64(cur.Val)*mult
		mult *= 10
		cur = cur.Next
	}
	return val
}

func EncodeNumber(i int64) *ListNode {
	dummy := &ListNode{}
	current := dummy
	for i > 0 || current == dummy {
		digit := i % 10
		i = i / 10
		newNode := &ListNode{
			Val: int(digit),
		}
		current.Next = newNode
		current = current.Next
	}

	return dummy.Next
}