type TimeMap struct {
	data map[string][]DataWithTime
}

type DataWithTime struct {
	data string
	time int
}

func Constructor() TimeMap {
	return TimeMap{
		data: map[string][]DataWithTime{},
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	data := DataWithTime{data:value, time:timestamp}
	this.data[key]= append(this.data[key], data)
}

func (this *TimeMap) Get(key string, timestamp int) string {
	values, ok := this.data[key]
	if !ok {
		return ""
	}
	if values[0].time > timestamp {
		return ""
	}
	// Perform binary search on times
	start:= 0
	end := len(values)-1
	var found DataWithTime
	for {
		if end<=start {
			if values[start].time <= timestamp {
				found = values[start]
			}
			break
		}
		middle := start+(end-start)/2
		if values[middle].time == timestamp {
			return values[middle].data
		} else if values[middle].time > timestamp {
			end = middle-1
		} else {
			found = values[middle]
			start = middle+1
		}
	}


	return found.data
}
