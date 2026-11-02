/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	nodeIndexes := map[*Node]int{}
	current := head
	i := 0
	dummy := &Node{}
	previous := dummy
	var nodes []*Node
	for current != nil {
		// copy linked list
		newNode := &Node{Val:current.Val}
		previous.Next = newNode
		previous = newNode
		nodes = append(nodes, newNode)

		// Get info for future Random Indexes
		nodeIndexes[current] = i
		i++
		current = current.Next
	}

	i = 0
	current = head
	for current != nil {
		if current.Random != nil {
			indexToPointTo := nodeIndexes[current.Random]
			nodes[i].Random = nodes[indexToPointTo]
		}
		current=current.Next
		i++
	}

	return dummy.Next
}
