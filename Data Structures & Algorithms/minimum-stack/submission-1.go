type MinStack struct {
	data []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	this.data = append(this.data, val)
}

func (this *MinStack) Pop() {
	this.data = this.data[:len(this.data)-1]
}

func (this *MinStack) Top() int {
	return this.data[len(this.data)-1]
}

func (this *MinStack) GetMin() int {
	min := math.MaxInt
	for _, val := range this.data {
		if val< min {
			min = val
		}
	}
	return min
}
