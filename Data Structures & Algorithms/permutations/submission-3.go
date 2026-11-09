
func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{[]int{}}
	}
	results := [][]int{make([]int, 0, len(nums))} 

	for _, num := range nums {
		var newResults [][]int
		for _, current := range results {
			for i := 0; i < len(results[0])+1; i++ {
				var next []int
				// if i == len(results[0]) {
				// 	next = append(current, num)
				// } else {
					next = make([]int, len(results[0])+1, len(nums))
					next[i] = num
					copy(next[0:i], current[0:i])
					copy(next[i+1:], current[i:])
				// }
				newResults = append(newResults, next)
			}
		}
		results = newResults
	}

	return results
}
