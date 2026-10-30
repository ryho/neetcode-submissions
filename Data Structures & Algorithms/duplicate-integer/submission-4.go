func hasDuplicate(nums []int) bool {
    check := map[int]struct{}{}
	for _, num := range nums {
		if _, ok := check[num]; ok {
			return true
		}
		check[num] = struct{}{}
	}
	return false
}
