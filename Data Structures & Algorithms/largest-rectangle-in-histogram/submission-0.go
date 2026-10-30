func largestRectangleArea(heights []int) int {
	// height/index
	var stack [][2]int
	var maxArea int
	for i, currentHeight := range heights {
		if len(stack) == 0 {
			stack = append(stack, [2]int{currentHeight, i})
		} else {
			peek := stack[len(stack)-1]
			lastIndex := i
			for ;len(stack)>0;{
				peek = stack[len(stack)-1]
				peekHeight := peek[0]
				peekIndex:= peek[1]
				if peekHeight <= currentHeight {
					break
				}
				lastIndex = peekIndex
				// If previous is taller, figure out that box
				area := (i - peekIndex)*peekHeight
				if area > maxArea {
					maxArea = area
				}
				stack = stack[:len(stack)-1]
				
			}
			// new box starts with the last one
			// skip if box has the same height
			if peek[0]!= currentHeight {
				stack = append(stack, [2]int{currentHeight,lastIndex})
			}
		}
	}
	for _,box := range stack {
		// all boxes in the stack go all the way to the end
		area := (len(heights)-box[1])*box[0]
		if area > maxArea {
			maxArea = area
		}
	}
	return maxArea
}