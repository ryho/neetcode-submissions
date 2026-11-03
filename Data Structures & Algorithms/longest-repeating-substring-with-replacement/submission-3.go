func characterReplacement(s string, k int) int {
	if len(s)==1 {
		return 1
	}
	var left, right, maxLength int
	counts := map[byte]int{}
	counts[s[0]]=1
	for i:=1;i<len(s);i++ {
		right=i
		counts[s[right]]++
		maxCount,otherCount,_ := asdf(counts)
		// fmt.Printf("maxCount %v otherCount %v counts %v\n", maxCount, otherCount,counts)
		for otherCount > k{
			// move left
			counts[s[left]]--
			left++
			maxCount,otherCount,_ = asdf(counts)
		}
		if maxCount + otherCount > maxLength {
			maxLength = maxCount+otherCount
		}
	}
	return maxLength
}

func asdf(input map[byte]int) (int,int,byte){
	maxVal := 0
	var maxChar byte
	other := 0
	for k,v := range input {
		if maxVal == 0 {
			maxVal = v
			maxChar = k
		}else {
			if v>maxVal {
				other+=maxVal
				maxVal = v
				maxChar = k
			} else {
				other+=v
			}
		}
	}
	return maxVal,other, maxChar
}