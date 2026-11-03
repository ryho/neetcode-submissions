func minWindow(s string, t string) string {
    tMap := map[rune]int{}
	for _,c := range t {
		tMap[c]++
	}
	sMap := map[rune]int{}
	var shortestSubString string
	left := 0
	neededCounts:= len(tMap)
	haveCounts := 0
	for right, c := range s {
		sMap[c]++
		if sMap[c] == tMap[c] && tMap[c]>0{
			haveCounts++
		}
		for haveCounts==neededCounts {
			if len(shortestSubString) == 0 || 
			len(shortestSubString) > (right-left+1) {
				shortestSubString = s[left:right+1]
			}
			lChar := rune(s[left])
			sMap[lChar]--
			if sMap[lChar]==0 {
				delete(sMap, lChar)
			}
			if sMap[lChar]==tMap[lChar]-1 {
				haveCounts--
			}
			left++
		}
	}

	return shortestSubString
}

func mapContains(tMap,sMap map[rune]int) bool {
	for k,v:= range tMap {
		if sMap[k]<v {
			return false
		}
	}
	return true
}
