package main

import "fmt"

func mergeIntervals(intervals []Interval) []Interval {
	// Replace this placeholder return statement with your code
	var output []Interval
	for i, interval := range intervals {
		if i == 0 {
			output = append(output, interval)
			continue
		}
		last := output[len(output)-1]
		if interval.start <= last.end {
			if interval.end > last.end {
				last.end = interval.end
			}
			output[len(output)-1] = last
		} else {
			output = append(output, interval)
		}

	}
	return output
}

// todobetul check
func insertInterval(existingIntervals []Interval, newInterval Interval) []Interval {
	// Write your code here
	// Your code will replace this placeholder return statement
	var output []Interval
	isprocessed := false
	for i, interval := range existingIntervals {
		if i == 0 {
			output = append(output, interval)
			continue
		}
		last := output[len(output)-1]

		if !isprocessed {
			if newInterval.start <= last.end {
				if newInterval.start < last.start {
					last.start = newInterval.start
				}

				if newInterval.end > last.end {
					last.end = newInterval.end
				}
				output[len(output)-1] = last
				isprocessed = true
			} else if newInterval.end < interval.start {
				output = append(output, newInterval)
				last = newInterval
				isprocessed = true
			}
		}

		if interval.start <= last.end {
			if interval.end > last.end {
				last.end = interval.end
			}
			output[len(output)-1] = last
		} else {
			output = append(output, interval)
		}

		if !isprocessed && i == len(existingIntervals)-1 {
			output = append(output, newInterval)
		}
	}
	return output
}

func main() {
	// Test your function
	intervals := []Interval{newInterval(1, 4), newInterval(2, 5), newInterval(7, 9)}
	fmt.Println("Merged intervals: ", mergeIntervals(intervals))

	intervals = []Interval{newInterval(2, 4), newInterval(5, 9), newInterval(6, 7)}
	fmt.Println("Merged intervals: ", mergeIntervals(intervals))

	intervals = []Interval{newInterval(1, 3), newInterval(5, 7), newInterval(8, 12)}
	fmt.Println("insertInterval intervals: ", insertInterval(intervals, newInterval(4, 6)))

	intervals = []Interval{newInterval(1, 2), newInterval(3, 4), newInterval(5, 8), newInterval(9, 15)}
	fmt.Println("insertInterval intervals: ", insertInterval(intervals, newInterval(16, 17)))

}
