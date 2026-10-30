import "slices"

func minEatingSpeed(piles []int, timeLimit int) int {
	minSpeed := 1
	maxSpeed := slices.Max(piles)
	for {
		var totalHours int
		if maxSpeed <= minSpeed {
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
