

func carFleet(target int, position []int, speed []int) int {
	if len(position) == 0 {
		return 0
	}

	cars := Cars([2][]int{position,speed})
	sort.Sort(cars)

	var currentfleetCount, countOfFleets int
	var currentTime float32
	for i, p := range position {
		time := float32(target-p)/float32(speed[i])
		if time <= currentTime {
			currentfleetCount++
		} else {
			countOfFleets++
			currentTime = time
			currentfleetCount = 1
		}
	}
	return countOfFleets
}
// position/speed

type Cars [2][]int

// 3. Implement sort.Interface methods
func (p Cars) Len() int           { return len(p[0]) }
func (p Cars) Less(i, j int) bool { return p[0][i] > p[0][j] }

// The Swap function modifies the elements in-place using Go's assignment syntax
func (p Cars) Swap(i, j int)      {
	p[0][i], p[0][j] = p[0][j], p[0][i]
	p[1][i], p[1][j] = p[1][j], p[1][i]
}
