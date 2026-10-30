

func carFleet(target int, position []int, speed []int) int {
	if len(position) == 0 {
		return 0
	}
	// // position/speed
	// cars := make([][2]int, len(position))
	// for i, p := range position {
	// 	cars[i] = [2]int{p,speed[i]}
	// }
	// // Sort high position to low
	// slices.SortFunc(cars, func(a, b [2]int) int {
	// 	return b[0]-a[0]
	// })
	var cars Cars
	cars = [2][]int{position,speed}
	sort.Sort(cars)
	var currentfleetCount, countOfFleets int
	var currentTime float32
	fmt.Printf("positions %v, times %v, target %v\n", position, speed, target)
	for i, p := range position {
		time := float32(target-p)/float32(speed[i])
		// fmt.Printf("target %v, p %v,  speed[i] %v, time %v, currentTime %v\n", target, p, speed[i], time, currentTime)
		if i == 0 {
			currentTime = time
			currentfleetCount++
			continue
		}
		if time <= currentTime {
			currentfleetCount++
		} else {
			countOfFleets++
			currentTime = time
			currentfleetCount = 1
		}
	}
	countOfFleets++
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
