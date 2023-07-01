package main

import (
	"fmt"
	"testing"
	"testing/quick"
)

func TestSolution(t *testing.T) {
	input := []int{9, 3, 9, 3, 7}
	expected := 7
	actual := SolutionSort(input)
	if actual != expected {
		t.Errorf("Test failed, expected: '%d', got:  '%d'", expected, actual)
	}
}

func TestFuzzSolutionSort(t *testing.T) {
	f := func(a []int) bool {
		fmt.Print("input array: ", a, "\n")
		result := SolutionSort(a)
		return result == Solutionxor(a)
	}

	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

//{-8034843474017254924, -5777498374867451097, -2289411683948202627, -947120291822621557, 2049642480346318233, 2120283203010594819, 4103990856394083237, 4527516735047530362, 8987734566510693546}
