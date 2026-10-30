func search(nums []int, target int) int {
	left := 0
	right := len(nums)-1
	for {
		if right<=left {
			if nums[left] == target{
				return left
			} else {
				return -1
			}
		}
		middle := left+(right-left)/2
		if nums[middle]==target {
			return middle
		} else if nums[middle]>target {
			right = middle-1
		} else {
			left = middle+1
		}
	}
}
