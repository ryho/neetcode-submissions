import "slices"
func topKFrequent(nums []int, k int) []int {
	data := map[int]int{}
	var filteredNums []int
	for _, n := range nums {
		if _, ok := data[n];!ok {
			filteredNums = append(filteredNums, n)
		}
		data[n]+=1
	}
	slices.SortFunc(filteredNums,func(a,b int) int{
		return data[b] - data[a]
	})
	return filteredNums[:k]
}
