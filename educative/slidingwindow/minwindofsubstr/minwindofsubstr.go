package main

import "fmt"

func minWindow(str1 string, str2 string) string {
	// your code will replace this placeholder return statement
	minst := str1

	ws, we := iterateString(str1, str2, 0)
	if we == 0 {
		return ""
	}
	if len(minst) > we-ws+1 {
		minst = str1[ws : we+1]
	}
	ws++
	for ws < len(str1) && we < len(str1)-1 {
		ws, we = iterateString(str1, str2, ws)
		if we == 0 {
			break
		}
		if len(minst) > we-ws+1 {
			minst = str1[ws : we+1]
		}
		ws++
	}
	return minst
}

func iterateString(str1 string, str2 string, st1 int) (int, int) {
	// your code will replace this placeholder return statement
	i1 := st1
	i2 := 0
	ws := st1
	we := 0

	for i1 < len(str1) && i2 < len(str2) {
		if str1[i1] == str2[i2] {
			if i2 == len(str2)-1 {
				we = i1
				break
			}
			i2++
		}
		i1++
	}

	if we == 0 {
		return ws, we
	}

	for i1 >= st1 && i2 >= 0 {
		if str1[i1] == str2[i2] {
			if i2 == 0 {
				ws = i1
				break
			}
			i2--
		}
		i1--
	}

	return ws, we
}

func main() {
	fmt.Println(minWindow("this is a test string", "tist"))
	fmt.Println(minWindow("abcdebdde", "tist"))
	fmt.Println(minWindow("abcdebde", "bde"))
}
