import "slices"

func minEatingSpeed(piles []int, timeLimit int) int {
	// timeLimit-=1
	slices.Sort(piles)
	minSpeed := 1
	maxSpeed := piles[len(piles)-1]
	for {
		var totalHours int
		if maxSpeed <= minSpeed {
			fmt.Printf("returnign min max %v min %v\n", maxSpeed, minSpeed)
			return minSpeed
		}

		currentSpeed := minSpeed + (maxSpeed-minSpeed)/2
		for _, pile := range piles {
			time := pile/currentSpeed
			if pile%currentSpeed > 0 {
				time++
			}
			totalHours += time
		}
		fmt.Printf("totalHours %v timeLimit %v currentSpeed %v max %v min %v\n", totalHours, timeLimit, currentSpeed, maxSpeed, minSpeed)
		if totalHours == timeLimit {
			maxSpeed = currentSpeed
		} else if totalHours > timeLimit {
			minSpeed = currentSpeed + 1
		} else if totalHours < timeLimit {
			maxSpeed = currentSpeed
		}
	}
	return 0
}
