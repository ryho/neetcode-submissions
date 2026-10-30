type DynamicArray struct {
	data []int
	capacity, length int
}

func NewDynamicArray(capacity int) *DynamicArray {
	return &DynamicArray{
		data: make([]int,capacity),
		capacity: capacity,
	}
}

func (da *DynamicArray) Get(i int) int {
	return da.data[i]
}

func (da *DynamicArray) Set(i int, n int) {
	da.data[i] = n
}

func (da *DynamicArray) Pushback(n int) {
	if da.length == da.capacity {
		da.resize()
	}
	da.data[da.length]=n
	da.length+=1
}

func (da *DynamicArray) Popback() int {
	da.length -=1
	return da.data[da.length]
}

func (da *DynamicArray) resize() {
	newData := make([]int,da.capacity*2)
	for i,x := range da.data {
		newData[i] = x
	}
	da.data = newData
	da.capacity *=2
}

func (da *DynamicArray) GetSize() int {
	return da.length
}

func (da *DynamicArray) GetCapacity() int {
	return da.capacity
}
