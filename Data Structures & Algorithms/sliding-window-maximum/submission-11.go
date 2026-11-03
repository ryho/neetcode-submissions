import "slices"
// An IntHeap is a max-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] } // > makes it a max-heap
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func maxSlidingWindow(nums []int, k int) []int {
    
listToLog := []int{1,-1}

	h := &IntHeap{}
	heap.Init(h)
	var maxes []int

	for i, num := range nums {
		heap.Push(h,num)
		if slices.Equal(nums,listToLog){fmt.Printf("h %v\n",h)}
		if i >= k-1 {
			max := heap.Pop(h).(int)
			
			if slices.Equal(nums,listToLog){fmt.Printf("popped %v\n",max)}
			maxes = append(maxes, max)

			if i == len(nums)-1{
				break
			}
			for {
				var found bool
				if slices.Equal(nums,listToLog){fmt.Printf("i %v, k %v, i-k+1 %v, i+1 %v\n", i, k , i-k+1+1,i+1)}
				if slices.Equal(nums,listToLog){fmt.Printf("slicedNums %v\n",nums[i-k+1+1:i+1])}
				for _, num := range nums[i-k+1+1:i+1] {
					if max == num {
						found = true
						break
					}
				}
				if found {
					break
				}
				if h.Len() == 0 {
					break
				}
				max = heap.Pop(h).(int)
			}
			if k > 1 {
				heap.Push(h,max)
			}
		}
	}
	return maxes
}
