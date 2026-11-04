
func trap(height []int) int {
	if len(height)<=2{
		return 0
	}
	left := 0
	right := len(height)-1
	totalTrapped := 0
	leftMax := height[0]
	rightMax := height[right]

	for right-1 != left {
		if height[left+1] >= leftMax {
			left++
			leftMax = height[left]
		} else if height[left+1] < rightMax && leftMax <= rightMax {
			left++
			trapped := leftMax-height[left]
			totalTrapped += trapped
		} else if height[right-1] > rightMax {
			right--	
			rightMax = height[right]
		} else if height[right-1] < leftMax && rightMax <= leftMax {
			right--
			trapped := rightMax-height[right]
			totalTrapped += trapped
		}
	}
	return totalTrapped
}
