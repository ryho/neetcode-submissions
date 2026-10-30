

type Solution struct{}


func (s *Solution) Encode(strs []string) string {
	if len(strs) ==0 {
		return ""
	}
	allStringsEmpty := true
	for i, str := range strs {
		if len(str) > 0 {
			allStringsEmpty = false
		}
		var wrapped bool
		if strings.Contains(str,"\""){
			str = "\""+strings.ReplaceAll(str,"\"","\"\"")+"\""
			strs[i] = str
			wrapped = true
		}
		if !wrapped && strings.Contains(str,",") {
			str = "\""+str+"\""
			strs[i] = str
		}
	}
	result := strings.Join(strs, ",")
	if allStringsEmpty {
		result +=","
	}
	fmt.Println(result)
	return result
}

func (s *Solution) Decode(encoded string) []string {
	if encoded == "" {
		return []string{}
	}
	var nextStr string
	newCell := true
	var skippedMiddleQuote bool
	var startedWithQuote bool
	var strs []string
	for i := range encoded {
		if newCell {
			if encoded[i] =='"' {
				startedWithQuote = true
			} else if encoded[i] ==',' {
				strs = append(strs, "")
				continue
			} else {
				startedWithQuote = false
				nextStr = encoded[i:i+1]
			}
			newCell = false
			skippedMiddleQuote = false
		} else {
			if encoded[i] =='"' {
				if skippedMiddleQuote {
					skippedMiddleQuote = false
					continue
				} else {
					nextStr += encoded[i:i+1]
					skippedMiddleQuote = true
					continue
				}
			} else if encoded[i] == ','{
				if skippedMiddleQuote || !startedWithQuote{
					if startedWithQuote {
						nextStr = nextStr[:len(nextStr)-1]
					}
					// this shoudld be the end of the word
					strs = append(strs, nextStr)
					nextStr = ""
					newCell = true
				} else {
					// this is a comma to insert
					nextStr += encoded[i:i+1]
				}
			} else {
				nextStr += encoded[i:i+1]
			}
		}
	}
	if startedWithQuote {
		nextStr = nextStr[:len(nextStr)-1]
	}
	strs = append(strs, nextStr)
	allStringsEmpty := true
	for _, str := range strs {
		if len(str) > 0 {
			allStringsEmpty = false
			break
		}
	}
	if allStringsEmpty {
		strs = strs[:len(strs)-1]
	}

	return strs
}
