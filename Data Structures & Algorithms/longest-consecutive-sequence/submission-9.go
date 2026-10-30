import "slices"

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	
	slices.Sort(nums)
	lastNum := nums[0]
	length := 1
	maxLength := 1
	for _, num := range nums {
		if num == lastNum {
			continue
		}
		if num == lastNum + 1 {
			length++
		} else {
			if length > maxLength {
				maxLength = length
			}
			length = 1
		}
		lastNum = num
	}
	if length > maxLength {
		maxLength = length
	}
	return maxLength
}
