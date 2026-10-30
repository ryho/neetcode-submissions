// 4:42
func maxArea(heights []int) int {
	i:= 0
	j:= len(heights)-1
	var maxWater int
	for {
		iVal := heights[i]
		jVal := heights[j]
		height := min(iVal,jVal)
		waterVol := height*(j-i)
		if waterVol > maxWater {
			maxWater = waterVol
		}
		if iVal<=jVal{
			i++
		} else {
			j--
		}
		if i == j {
			break
		}
	}
	return maxWater
}
