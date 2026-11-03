
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
    h := &IntHeap{}
	heap.Init(h)
	var maxes []int
	for i, num := range nums {
		heap.Push(h,[2]int{num,i})
		if i >= k-1 {
			maxes = append(maxes, peek(h)[0])
			for h.Len() > 0 && peek(h)[1] <= i-k+1 {
				heap.Pop(h)
			}
		}
	}
	return maxes
}

func peek(h *IntHeap) [2]int{
	return (*h)[0]
}
