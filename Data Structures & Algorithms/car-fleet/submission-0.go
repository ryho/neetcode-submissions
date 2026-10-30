import "slices"

func carFleet(target int, position []int, speed []int) int {
	if len(position) == 0 {
		return 0
	}
	// position/speed
	cars := make([][2]int, len(position))
	for i, p := range position {
		cars[i] = [2]int{p,speed[i]}
	}
	// Sort high position to low
	slices.SortFunc(cars, func(a, b [2]int) int {
		return b[0]-a[0]
	})
	var currentfleetCount, countOfFleets int
	var currentTime float32
	for i, car := range cars {
		time := float32(target-car[0])/float32(car[1])
		// fmt.Printf("target %v, car %v, time %v, currentTime %v\n", target, car, time, currentTime)
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
