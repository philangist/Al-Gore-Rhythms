package sorting

import (
    "fmt"
    "reflect"
    "testing"
)

type mergeSortTestCases struct {
    tag            string
    unsorted_input []int
    sorted_output  []int
}

func TestMergeSort(t *testing.T) {
    fmt.Println("Running TestMergeSort...")
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
            unsorted_input: []int{8,7,6,5, 4, 3, 2, 1},
            sorted_output:  []int{1, 2, 3, 4, 5,6,7,8},
        },
        {
            tag:            "Case 5",
            unsorted_input: []int{-1, 0, -999, 0xFFFF, 14, 42},
            sorted_output:  []int{-999, -1, 0, 14, 42, 0xFFFF},
        },
    }

    for _, testCase := range testCases {
        fmt.Println("Generating sort for: ", testCase.unsorted_input)
        observed_output := MergeSort(testCase.unsorted_input)
        if !reflect.DeepEqual(testCase.sorted_output, observed_output) {
            t.Errorf(
                "%s: Sort of %v was incorrect.\nExpected: %v\nReceived: %v",
                testCase.tag, testCase.unsorted_input, testCase.sorted_output, observed_output,
            )
        }
    }
}
