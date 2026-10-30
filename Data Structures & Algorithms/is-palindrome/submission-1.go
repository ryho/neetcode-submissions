func isPalindrome(s string) bool {
	j := len(s)-1
	for i, c := range s {
		if !((c >='a' && c <='z') || (c >='A' && c <='Z') || (c >='0' && c <='9') ){
			continue
		}
		if c >='A' && c <='Z' {
			c=c-'A'+'a'
		}
		for {
			if j < 0 {
				return true
			}
			c2 := rune(s[j])
			if !( (c2 >='a' && c2 <='z') || (c2 >='A' && c2 <='Z') || (c2 >='0' && c2 <='9')) {
				j--
				continue
			}
			if i >= j {
				return true
			}

			if c2 >='A' && c2 <='Z' {
				c2=c2-'A'+'a'
			}
			if c!=c2 {
				fmt.Printf("error c %c c2 %c i %v j %v\n", c, c2, i, j)
				return false
			} else {

				fmt.Printf("match c %c c2 %c i %v j %v\n", c, c2, i, j)
			}
			j--
			break
		}

	}
	return true
}
