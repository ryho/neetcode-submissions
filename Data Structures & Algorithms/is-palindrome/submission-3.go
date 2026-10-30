func isPalindrome(s string) bool {
	j := len(s)-1
	for i, c := range s {
		if !(isAlphaNumeric(c)) {
			continue
		}
		c = lowerCase(c)
		for {
			if i >= j {
				// Once they meet or overlap, we are done
				return true
			}
			c2 := rune(s[j])
			if !(isAlphaNumeric(c2)) {
				j--
				continue
			}
			c2 = lowerCase(c2)
			if c!=c2 {
				return false
			}
			j--
			break
		}
	}
	return true
}

func isAlphaNumeric(c rune) bool {
	return (c >='a' && c <='z') || (c >='A' && c <='Z') || (c >='0' && c <='9')
}

func lowerCase(c rune) rune{
	if c >='A' && c <='Z' {
		return c-'A'+'a'
	}
	return c
}
