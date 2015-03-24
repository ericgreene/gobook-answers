package main

import (
	"fmt"
)

// Create a function that accepts an []int and returns an []int which is a copy of the given
// []int but with duplicates removed.
func main() {
	input := []int{9, 1, 9, 5, 4, 4, 2, 1, 5, 4, 8, 8, 4, 3, 6, 9, 5, 7, 5}
	output := UniqueInts(input)
	fmt.Println(output)
}

func UniqueInts(input []int) []int {
	output := []int{}
	var found bool

	for _, val1 := range input {
		found = false
		for _, val2 := range output {
			if val1 == val2 {
				found = true
			}
		}
		if !found {
			output = append(output, val1)
		}
	}
	return output
}
