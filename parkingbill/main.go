package main

import (
	"fmt"
	"time"
)

const (
	fix         = 2
	ilksaat     = 3
	sonrakisaat = 4
)

func Solution(E string, L string) int {
	formatter := "15:04"
	etime, _ := time.Parse(formatter, E)
	ltime, _ := time.Parse(formatter, L)

	cost := fix
	if ltime.Hour() > etime.Hour() {
		cost += ilksaat
		cost += sonrakisaat * ((ltime.Hour() - etime.Hour()) - 1)
		if ltime.Minute() > etime.Minute() {
			cost += sonrakisaat
		}
	} else {
		if ltime.Minute() > etime.Minute() {
			cost += ilksaat
		}
	}

	return cost
}

func main() {
	fmt.Print(Solution("10:00", "13:21"))
}
