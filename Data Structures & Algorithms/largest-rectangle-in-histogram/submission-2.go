func largestRectangleArea(heights []int) int {
	heights = append(heights,0)
	// 0:height 1:index
	var stack [][2]int
	var maxArea int
	for i, currentHeight := range heights {
		if len(stack) == 0 {
			stack = append(stack, [2]int{currentHeight, i})
		} else {
			peek := stack[len(stack)-1]
			newBoxIndex := i
			for ;len(stack)>0;{
				peek = stack[len(stack)-1]
				peekHeight := peek[0]
				peekIndex:= peek[1]
				if peekHeight <= currentHeight {
					break
				}
				newBoxIndex = peekIndex
				// If previous is taller, figure out that box
				area := (i - peekIndex)*peekHeight
				if area > maxArea {
					maxArea = area
				}
				stack = stack[:len(stack)-1]
				
			}
			// new box starts with the last index we found
			// skip if box has the same height
			if peek[0]!= currentHeight {
				stack = append(stack, [2]int{currentHeight,newBoxIndex})
			}
		}
	}
	return maxArea
}