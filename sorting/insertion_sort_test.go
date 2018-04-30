package sorting

import (
	"fmt"
	"reflect"
	"testing"
)

type insertionSortTestCases struct {
	tag            string
	unsorted_input []int
	sorted_output  []int
}

func TestLinearSearchInsertionSort(t *testing.T) {
	fmt.Println("Running TestLinearSearchInsertionSort...")
	testCases := []insertionSortTestCases{
		{
			tag:            "Case 1",
			unsorted_input: []int{},
			sorted_output:  []int{},
		},
		{
			tag:            "Case 2",
			unsorted_input: []int{9},
			sorted_output:  []int{9},
		},
		{
			tag:            "Case 3",
			unsorted_input: []int{1, 2, 3, 4, 5},
			sorted_output:  []int{1, 2, 3, 4, 5},
		},
		{
			tag:            "Case 4",
			unsorted_input: []int{5, 4, 3, 2, 1},
			sorted_output:  []int{1, 2, 3, 4, 5},
		},
		{
			tag:            "Case 5",
			unsorted_input: []int{-1, 0, -999, 0xFFFF, 14, 42},
			sorted_output:  []int{-999, -1, 0, 14, 42, 0xFFFF},
		},
	}

	for _, testCase := range testCases {
		fmt.Println("Generation sort with linearSearchInsert for: ", testCase.unsorted_input)
		observed_output := InsertionSort(testCase.unsorted_input, linearSearchInsert)
		if !reflect.DeepEqual(testCase.sorted_output, observed_output) {
			t.Errorf(
				"%s: Sort of %v was incorrect.\nExpected: %v\nReceived: %v",
				testCase.tag, testCase.unsorted_input, testCase.sorted_output, observed_output,
			)
		}
	}
}

func TestBinarySearchInsertionSort(t *testing.T) {
	fmt.Println("Running TestBinarySearchInsertionSort...")
	testCases := []insertionSortTestCases{
		{
			tag:            "Case 1",
			unsorted_input: []int{},
			sorted_output:  []int{},
		},
		{
			tag:            "Case 2",
			unsorted_input: []int{9},
			sorted_output:  []int{9},
		},
		{
			tag:            "Case 3",
			unsorted_input: []int{1, 2, 3, 4, 5},
			sorted_output:  []int{1, 2, 3, 4, 5},
		},
		{
			tag:            "Case 4",
			unsorted_input: []int{5, 4, 3, 2, 1},
			sorted_output:  []int{1, 2, 3, 4, 5},
		},
		{
			tag:            "Case 5",
			unsorted_input: []int{-1, 0, -999, 0xFFFF, 14, 42},
			sorted_output:  []int{-999, -1, 0, 14, 42, 0xFFFF},
		},
	}

	for _, testCase := range testCases {
		fmt.Println("Generation sort with binarySearchInsert for: ", testCase.unsorted_input)
		observed_output := InsertionSort(testCase.unsorted_input, binarySearchInsert)
		if !reflect.DeepEqual(testCase.sorted_output, observed_output) {
			t.Errorf(
				"%s: Sort of %v was incorrect.\nExpected: %v\nReceived: %v",
				testCase.tag, testCase.unsorted_input, testCase.sorted_output, observed_output,
			)
		}
	}
}
