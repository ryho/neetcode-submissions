

func carFleet(target int, position []int, speed []int) int {
	if len(position) == 0 {
		return 0
	}

	// Sort in place
	cars := Cars([2][]int{position,speed})
	sort.Sort(cars)

	var currentfleetCount, countOfFleets int
	var currentTime float32
	for i, p := range position {
		time := float32(target-p)/float32(speed[i])
		if time <= currentTime {
			// car will join the fleet
			currentfleetCount++
		} else {
			// car is too slow, makes a new fleet
			countOfFleets++
			currentTime = time
			currentfleetCount = 1
		}
	}
	return countOfFleets
}

// 0:position,1:speed
type Cars [2][]int
func (p Cars) Len() int           { return len(p[0]) }
func (p Cars) Less(i, j int) bool { return p[0][i] > p[0][j] }
func (p Cars) Swap(i, j int)      {
	p[0][i], p[0][j] = p[0][j], p[0][i]
	p[1][i], p[1][j] = p[1][j], p[1][i]
}
