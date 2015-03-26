package main

import (
	"reflect"
	"testing"
)

func TestUniqueInts(t *testing.T) {
	input := []int{9, 1, 9, 5, 4, 4, 2, 1, 5, 4, 8, 8, 4, 3, 6, 9, 5, 7, 5}
	assertDeepEqual(t, UniqueInts1(input), []int{9, 1, 5, 4, 2, 8, 3, 6, 7}, "arrays not equal", 1)
	assertDeepEqual(t, UniqueInts2(input), []int{9, 1, 5, 4, 2, 8, 3, 6, 7}, "arrays not equal", 2)
}

func TestMake2D(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	output := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {10, 11, 12}, {13, 14, 15}, {16, 17, 18}, {19, 20, 0}}
	output2 := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}, {17, 18, 19, 20}}
	output3 := [][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13, 14, 15}, {16, 17, 18, 19, 20}}
	output4 := [][]int{{1, 2, 3, 4, 5, 6}, {7, 8, 9, 10, 11, 12}, {13, 14, 15, 16, 17, 18}, {19, 20, 0, 0, 0, 0}}

	assertDeepEqual(t, output, Make2D(input, 3), "results not equal", 4)
	assertDeepEqual(t, output2, Make2D(input, 4), "results not equal", 5)
	assertDeepEqual(t, output3, Make2D(input, 5), "results not equal", 6)
	assertDeepEqual(t, output4, Make2D(input, 6), "results not equal", 7)
}

func TestFlatten(t *testing.T) {
	irregularMatrix := [][]int{{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11},
		{12, 13, 14, 15},
		{16, 17, 18, 19, 20}}
	// output := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	output := []int{}
	for i := 0; i < 20; i++ {
		output = append(output, i+1)
	}
	assertDeepEqual(t, Flatten1(irregularMatrix), output, "slices are not equal", 3)
}

func BenchmarkUniqueInts1(b *testing.B) {
	// run the Fib function b.N times
	input := []int{9, 1, 9, 5, 4, 4, 2, 1, 5, 4, 8, 8, 4, 3, 6, 9, 5, 7, 5}
	for n := 0; n < b.N; n++ {
		UniqueInts1(input)
	}
}

func BenchmarkUniqueInts2(b *testing.B) {
	// run the Fib function b.N times
	input := []int{9, 1, 9, 5, 4, 4, 2, 1, 5, 4, 8, 8, 4, 3, 6, 9, 5, 7, 5}
	for n := 0; n < b.N; n++ {
		UniqueInts2(input)
	}
}

func BenchmarkFlatten1(b *testing.B) {
	irregularMatrix := [][]int{{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11},
		{12, 13, 14, 15},
		{16, 17, 18, 19, 20}}
	for n := 0; n < b.N; n++ {
		Flatten1(irregularMatrix)
	}
}

func BenchmarkFlatten2(b *testing.B) {
	irregularMatrix := [][]int{{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11},
		{12, 13, 14, 15},
		{16, 17, 18, 19, 20}}
	for n := 0; n < b.N; n++ {
		Flatten2(irregularMatrix)
	}
}

func assertDeepEqual(t *testing.T, a1, a2 interface{}, message string, id int) {
	if !reflect.DeepEqual(a1, a2) {
		t.Errorf("#%d: %v does not equal %v", id, a1, a2)
	}
}
