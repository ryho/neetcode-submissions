import "slices"
// An IntHeap is a max-heap of ints.
// 0:value 1:previous index
type IntHeap [][2]int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i][0] > h[j][0] } // > makes it a max-heap
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func maxSlidingWindow(nums []int, k int) []int {
	listToPrint := []int{1,-1}
    h := &IntHeap{}
	heap.Init(h)
	var maxes []int

	for i, num := range nums {
		
		heap.Push(h,[2]int{num,i})
		if i >= k-1 {
			if slices.Equal(nums, listToPrint) {fmt.Printf("h %v\n", h)}
			maxes = append(maxes, (*h)[0][0])
			for h.Len() > 0 && (*h)[0][1] <= i-k+1 {
				heap.Pop(h)
			}
		}
	}
	return maxes
}
