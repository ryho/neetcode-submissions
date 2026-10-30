type MinStack struct {
	stack []int
	min int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	if len(this.stack) == 0 {
		this.min = val
		this.stack = append(this.stack,0)
	} else {
		diff := val-this.min
		this.stack = append(this.stack, diff)
		this.min = min(this.min,val)
	}
}
// 1 2 0
// 0 1
// 1
// 0 1 -1
// 0

func (this *MinStack) Pop() {
	// fmt.Printf("pop %v %v\n", this.stack, this.min)
	top := this.stack[len(this.stack)-1]
	if top < 0 {
		this.min = this.min - top
	}
	this.stack = this.stack[:len(this.stack)-1]
	// fmt.Printf("pop %v %v\n", this.stack, this.min)
}

func (this *MinStack) Top() int {
	top:= this.stack[len(this.stack)-1]
	if top < 0 {
		return this.min
	}
	return top + this.min
}

func (this *MinStack) GetMin() int {
	return this.min
}
/*
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
*/

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