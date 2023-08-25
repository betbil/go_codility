package main

import (
	"fmt"
	"sort"
	"strings"
)

// TimeStamp data structure
type TimeStamp struct {
	// write your code here
	valuem   map[string]map[int]string //key - (timeStamp-value)
	keytimem map[string]int
}

// TimeStamp constructor
func Constructor() TimeStamp {
	// Replace this placeholder return statement with your code
	t := TimeStamp{}
	t.valuem = make(map[string]map[int]string)
	t.keytimem = make(map[string]int)
	return t
}

// Set TimeStamp data variables
func (this *TimeStamp) SetValue(key string, value string, timeStamp int) bool {
	// Replace this placeholder return statement with your code
	time, exits := this.keytimem[key]
	if exits && time > timeStamp {
		return false
	}
	if !exits {
		this.valuem[key] = make(map[int]string)
	}
	this.keytimem[key] = timeStamp
	this.valuem[key][timeStamp] = value
	return false
}

// Get the value for the given key and timestamp
func (this *TimeStamp) GetValue(key string, timestamp int) string {
	// Replace this placeholder return statement with your code
	time, ok := this.keytimem[key]
	if !ok {
		return ""
	}

	if time < timestamp {
		timestamp = time
	}

	val, ok := this.valuem[key][timestamp]
	if ok {
		return val
	}

	return ""
}

type valuetimestamp struct {
	value     string
	timestamp int
}

// ---
// TimeStamp data structure
type TimeStamp2 struct {
	// write your code here
	valuem map[string][]valuetimestamp
}

// TimeStamp constructor
func Constructor2() TimeStamp2 {
	// Replace this placeholder return statement with your code
	t := TimeStamp2{}
	t.valuem = make(map[string][]valuetimestamp)
	return t
}

// Set TimeStamp data variables
func (this *TimeStamp2) SetValue(key string, value string, timeStamp int) bool {
	// Replace this placeholder return statement with your code
	list, exits := this.valuem[key]
	if exits && list[len(list)-1].timestamp > timeStamp {
		return false
	}
	if !exits {
		this.valuem[key] = make([]valuetimestamp, 0)
	}
	this.valuem[key] = append(this.valuem[key], valuetimestamp{value, timeStamp})
	return true
}

// Get the value for the given key and timestamp
func (this *TimeStamp2) GetValue(key string, timestamp int) string {
	// Replace this placeholder return statement with your code
	list, exits := this.valuem[key]
	if !exits {
		return ""
	}
	time2find := list[len(list)-1].timestamp
	if time2find >= timestamp {
		time2find = timestamp
	}

	//find
	l := 0
	h := len(list) - 1
	for l < h {
		mid := l + (h-l)/2
		if list[mid].timestamp == time2find {
			return list[mid].value
		} else if list[mid].timestamp < time2find {
			l = mid + 1
			if list[l].timestamp > time2find {
				return list[mid].value
			}
		} else {
			h = mid
		}
	}

	return list[l].value

}

// -------
// TimeStamp data structure
type TimeStamp3 struct {
	valuesMap    map[string][]string
	timestampMap map[string][]int
}

// TimeStamp constructor
func Constructor3() TimeStamp3 {
	return TimeStamp3{
		valuesMap:    make(map[string][]string), // valuesMap hash map
		timestampMap: make(map[string][]int),    // timeStamp hash map
	}
}

// Set TimeStamp data variables in function.
func (this *TimeStamp3) SetValue(key string, value string, timeStamp int) {
	if _, ok := this.valuesMap[key]; ok {
		if value != this.valuesMap[key][len(this.valuesMap[key])-1] {
			this.valuesMap[key] = append(this.valuesMap[key], value)
			this.timestampMap[key] = append(this.timestampMap[key], timeStamp)
		}
	} else {
		this.valuesMap[key] = append(this.valuesMap[key], value)
		this.timestampMap[key] = append(this.timestampMap[key], timeStamp)
	}
}

// Get the value for the given key and timestamp
func (this *TimeStamp3) GetValue(key string, timestamp int) string {
	if _, ok := this.valuesMap[key]; !ok {
		return ""
	} else {
		index := sort.Search(len(this.timestampMap[key]),
			func(index int) bool { return this.timestampMap[key][index] > timestamp }) - 1
		if index > -1 {
			return this.valuesMap[key][index]
		}
		return ""
	}
}

func main() {

	//["TimeStamp", "set_value", "set_value", "get_value", "get_value"] , [[], ["foo", "tan", 7], ["foo", "ban", 9], ["foo", 8], ["foo", 9]]
	t := Constructor2()
	t.SetValue("foo", "tan", 7)
	t.SetValue("foo", "ban", 9)
	fmt.Println(t.GetValue("foo", 8))
	fmt.Println(t.GetValue("foo", 9))

	//--------------------------------
	s := "abdeklm"
	aa := "c"

	for i, v := range s {
		fmt.Println(i, string(v))
	}

	indx1 := sort.Search(len(s), func(i int) bool {
		return s[i] >= aa[0]
	})
	fmt.Println("indx1", indx1)

	fmt.Println("strings.Compare(string(s[0]), aa)", strings.Compare(string(s[0]), aa))

	indx, found := sort.Find(len(s), func(i int) int {
		cmp := strings.Compare(aa, string(s[i]))
		fmt.Println("i", i, "s[i]", string(s[i]), "aa", aa, "cmp", cmp)
		return cmp
	})
	fmt.Println(indx, found)

}
