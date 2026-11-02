func findDuplicate(nums []int) int {
    slow := 0
	fast := 0
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if fast == slow {
			break
		}
	}
	slow2 := 0
	for {
		slow = nums[slow]
		slow2 = nums[slow2]
		if slow == slow2 {
			break
		}
	}
	return slow
}
