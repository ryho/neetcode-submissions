//3:03


func findMin(nums []int) int {
	if len(nums) ==1 {
		return nums[0]
	}
	fmt.Printf("%v\n", nums)
	left := 0
	right := len(nums)-1
	for {
		if left == right -1 {
			if nums[left] < nums[right] {
				return nums[left]
			}
			return nums[right]
		}
		middle := left + (right-left)/2
		if nums[middle] >= nums[right] {
			left = middle
		} else {
			right = middle
		}
	}
	return 0
}
