type MinStack struct {
	stack []int
	minStack []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	min := val
	if len(this.minStack) > 0 && this.GetMin()<min{
		min = this.GetMin()
	}
	this.minStack = append(this.minStack, min)
	this.stack = append(this.stack, val)
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}

/*
type MinStackNaive struct {
	data []int
}

func Constructor() MinStackNaive {
	return MinStackNaive{}
}

func (this *MinStackNaive) Push(val int) {
	this.data = append(this.data, val)
}

func (this *MinStackNaive) Pop() {
	this.data = this.data[:len(this.data)-1]
}

func (this *MinStackNaive) Top() int {
	return this.data[len(this.data)-1]
}

func (this *MinStackNaive) GetMin() int {
	min := math.MaxInt
	for _, val := range this.data {
		if val< min {
			min = val
		}
	}
	return min
}
*/