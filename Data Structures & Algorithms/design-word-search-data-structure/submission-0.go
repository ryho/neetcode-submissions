type WordDictionary struct {
    data [27]*WordDictionary
}

func Constructor() WordDictionary {
    return WordDictionary{
		data:[27]*WordDictionary{},
	}
}
const END_CHAR = 'z'-'a'+1

func (this *WordDictionary) AddWord(word string)  {
	if len(word) == 0 {
		this.data[END_CHAR] = &WordDictionary{}
		return
	}
    c := word[0]
	next := this.data[c-'a']
	if next == nil {
		next = &WordDictionary{
			data:[27]*WordDictionary{},
		}
		this.data[c-'a'] = next
	}
	next.AddWord(word[1:])
}

func (this *WordDictionary) Search(word string) bool {
	if len(word)==0 {
		return this.data[END_CHAR] != nil
	}
    c := word[0]
	if c == '.' {
		for _, next := range this.data {
			if next != nil {
				found := next.Search(word[1:])
				if found {
					return true
				}
			}
		}
		return false
	} else {
		next := this.data[c-'a']
		if next == nil {
			return false
		}
		return next.Search(word[1:])
	}
}
