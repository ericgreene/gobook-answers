package main

import (
	"fmt"
)

// Create a function that accepts an []int and returns an []int which is a copy of the given
// []int but with duplicates removed.
func main() {
	input := []int{9, 1, 9, 5, 4, 4, 2, 1, 5, 4, 8, 8, 4, 3, 6, 9, 5, 7, 5, 9, 1, 19, 15, 14, 14, 2, 1, 5, 4, 18, 18, 4, 3, 6, 9, 5, 7, 5}
	output := UniqueInts1(input)
	output2 := UniqueInts2(input)
	fmt.Println(output)
	fmt.Println(output2)
}

func Flatten1(input [][]int) []int {
	output := []int{}
	for _, v := range input {
		for _, v2 := range v {
			output = append(output, v2)
		}
	}
	return output
}

func Flatten2(matrix [][]int) []int {
	slice := make([]int, 0, len(matrix)+len(matrix[0]))
	for _, innerSlice := range matrix {
		for _, x := range innerSlice {
			slice = append(slice, x)
		}
	}
	return slice
}

func UniqueInts1(input []int) []int {
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

func UniqueInts2(input []int) []int {
	output := []int{}
	vals := make(map[int]bool)
	for _, x := range input {
		if _, found := vals[x]; !found {
			output = append(output, x)
			vals[x] = true
		}
	}
	return output
}

func Make2D(input []int, columns int) [][]int {
	output := [][]int{}
	innerSlice := make([]int, 0, columns)
	index := 0

	for index < len(input) {
		if (index + columns) < len(input) {
			innerSlice = input[index : index+columns]
			index = index + columns
			output = append(output, innerSlice)
		} else {
			lastSlice := input[index:]
			remainder := (index + columns) - len(input) - 1
			for i := 0; i <= remainder; i++ {
				lastSlice = append(lastSlice, 0)
			}
			output = append(output, lastSlice)
			break
		}
	}
	return output
}
