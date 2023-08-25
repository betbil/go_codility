package main

import "fmt"

type SnapshotArray struct {
	// Write your code here
	values  []int
	snaps   map[int][]int
	snapcnt int
}

// Constructor intializes the SnapshotArray type object
func Constructor(length int) SnapshotArray {
	// Your code will replace this placeholder return statement
	s := SnapshotArray{}
	s.values = make([]int, length)
	s.snaps = make(map[int][]int)
	s.snapcnt = 0
	return s
}

// SetValue sets the value at a given index idx to val.
func (s *SnapshotArray) SetValue(idx int, val int) {
	// Write your code here
	s.values[idx] = val
}

// Snapshot function takes no parameters and returns the snapid.
// snapid is the number of times that the snapshot() function was called minus 1.
func (s *SnapshotArray) Snapshot() int {
	// Replace this placeholder return statement with your code
	copya := make([]int, len(s.values))
	copy(copya, s.values)
	s.snaps[s.snapcnt] = copya
	s.snapcnt++
	return s.snapcnt - 1
}

// GetValue returns the value at the index idx with the given snapid.
func (s *SnapshotArray) GetValue(idx int, snapshotId int) int {
	// Replace this placeholder return statement with your code
	list, ok := s.snaps[snapshotId]
	if !ok {
		return -1
	}

	return list[idx]
}

func main() {
	// Your code here
	s := Constructor(3)
	s.SetValue(0, 5)
	s.SetValue(1, 6)
	s.SetValue(2, 7)
	fmt.Println(s.Snapshot())
	s.SetValue(0, 8)
	fmt.Println(s.Snapshot())
	s.SetValue(1, 9)
	fmt.Println(s.Snapshot())
	fmt.Println("----00000----")
	fmt.Println(s.GetValue(0, 0))
	fmt.Println(s.GetValue(1, 0))
	fmt.Println(s.GetValue(2, 0))
	fmt.Println("-----11111---")
	fmt.Println(s.GetValue(0, 1))
	fmt.Println(s.GetValue(1, 1))
	fmt.Println(s.GetValue(2, 1))
	fmt.Println("----22222----")
	fmt.Println(s.GetValue(0, 2))
	fmt.Println(s.GetValue(1, 2))
	fmt.Println(s.GetValue(2, 2))
	fmt.Println("----3333----")
	fmt.Println(s.GetValue(0, 3))
	fmt.Println(s.GetValue(1, 3))
	fmt.Println(s.GetValue(2, 3))
}
