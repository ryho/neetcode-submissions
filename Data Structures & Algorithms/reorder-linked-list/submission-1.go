/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }

1 2 3 4 5
1 2 3 4

12
345
543

15243

12
34
43

1423



 */

func reorderList(head *ListNode) {
	printAll(head)
	// figure out length of list
    end := head 
	lengthOfList := 1
	for end.Next != nil {
		end = end.Next
		lengthOfList++
	}
	if lengthOfList == 1 {
		return
	}

	// split list into two sections
	splitLength := (lengthOfList + 2 - 1) / 2
	cur := head
	for i := 0; i < splitLength-1; i++ {
		cur = cur.Next
	}
	secondList := cur.Next
	cur.Next = nil

	printAll(head)

	printAll(secondList)

	// reorder second section
	current := secondList
	var prev *ListNode
	for current.Next != nil {
		next := current.Next
		current.Next = prev
		prev=current
		current = next
	}
	current.Next = prev
	secondList = current

	printAll(secondList)

	// merge lists together, every other element
	newList := &ListNode{}
	firstList := head
	current = newList
	for secondList != nil {
		current.Next = firstList
		current = current.Next
		firstList=firstList.Next

		current.Next = secondList
		current = current.Next
		secondList = secondList.Next
	}
	if firstList != nil {
		current.Next = firstList
	}
}

func printAll(head *ListNode) {
	for head != nil{
		fmt.Printf("%v ", head.Val)
		head = head.Next
	}
	fmt.Printf("\n")
}
