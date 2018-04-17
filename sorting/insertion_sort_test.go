package sorting

import (
    "fmt"
    "reflect"
    "testing"
)

type insertionSortTestCases struct {
    unsorted_input []int
    sorted_output  []int
}

func TestLinearSearchInsertionSort(t *testing.T) {
    fmt.Println("Running TestLinearSearchInsertionSort...")
    testCases := []insertionSortTestCases{
        {
            unsorted_input: []int{},
            sorted_output:  []int{},
        },
        {
            unsorted_input: []int{9},
            sorted_output:  []int{9},
        },
        {
            unsorted_input: []int{1, 2, 3, 4, 5},
            sorted_output:  []int{1, 2, 3, 4, 5},
        },
        {
            unsorted_input: []int{5, 4, 3, 2, 1},
            sorted_output:  []int{1, 2, 3, 4, 5},
        },
        {
            unsorted_input: []int{-1, 0, -999, 0xFFFF, 14, 42},
            sorted_output:  []int{-999, -1, 0, 14, 42, 0xFFFF},
        },
    }

    for _, testCase := range testCases {
        fmt.Println("Running with linearSearchInsert for: ", testCase.unsorted_input)
        observed_output := InsertionSort(testCase.unsorted_input, linearSearchInsert)
        if !reflect.DeepEqual(testCase.sorted_output, observed_output) {
            t.Errorf(
                "Sort of %v was incorrect.\nExpected: %v\nReceived: %v",
                testCase.unsorted_input, testCase.sorted_output, observed_output,
            )
        }
    }
}

func TestBinarySearchInsertionSort(t *testing.T) {
    fmt.Println("Running TestBinarySearchInsertionSort...")
    testCases := []insertionSortTestCases{
        {
            unsorted_input: []int{},
            sorted_output:  []int{},
        },
        {
            unsorted_input: []int{9},
            sorted_output:  []int{9},
        },
        {
            unsorted_input: []int{1, 2, 3, 4, 5},
            sorted_output:  []int{1, 2, 3, 4, 5},
        },
        {
            unsorted_input: []int{5, 4, 3, 2, 1},
            sorted_output:  []int{1, 2, 3, 4, 5},
        },
        {
            unsorted_input: []int{-1, 0, -999, 0xFFFF, 14, 42},
            sorted_output:  []int{-999, -1, 0, 14, 42, 0xFFFF},
        },
    }

    for _, testCase := range testCases {
        fmt.Println("Running with binarySearchInsert for: ", testCase.unsorted_input)
        observed_output := InsertionSort(testCase.unsorted_input, binarySearchInsert)
        if !reflect.DeepEqual(testCase.sorted_output, observed_output) {
            t.Errorf(
                "Sort of %v was incorrect.\nExpected: %v\nReceived: %v",
                testCase.unsorted_input, testCase.sorted_output, observed_output,
            )
        }
    }
}