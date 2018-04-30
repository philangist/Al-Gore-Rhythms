package trees

import (
	"fmt"
	"reflect"
	"testing"
)

type binarySearchTestCases struct {
	tag          string
	inputValues  []int
	destination  int
	expectedPath []int
}

func TestBinarySearchTree(t *testing.T) {
	fmt.Println("Running TestBinarySearchTree...")
	testCases := []binarySearchTestCases{
		{
			tag:          "Case 1",
			inputValues:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
			destination:  1,
			expectedPath: []int{8, 4, 2, 1},
		},
		{
			tag:          "Case 2",
			inputValues:  []int{1, 2, 3},
			destination:  1,
			expectedPath: []int{2, 1},
		},
		{
			tag:          "Case 3",
			inputValues:  []int{852, 3270, 474, 3197, 1840, 3377, 539, 97, 2691, 3812, 1365, 260, 2241, 887, 1355},
			destination:  97,
			expectedPath: []int{1365, 539, 260, 97},
		},
		{
			tag:          "Case 4",
			inputValues:  []int{10, 9, 11},
			destination:  21,
			expectedPath: []int{},
		},
	}

	for _, testCase := range testCases {
		fmt.Println("Create binary search tree structure for ", testCase.inputValues)
		searchTree := NewSearchTree(testCase.inputValues)
		actualPath := ExtractPathToRoot(BreadthFirstSearch(searchTree, testCase.destination))
		if !reflect.DeepEqual(testCase.expectedPath, actualPath) {
			t.Errorf("%s: Binary search tree creation failed.\nExpected path to destination %d: %v\nObserved path:%v",
				testCase.tag, testCase.destination, testCase.expectedPath, actualPath,
			)
		}
	}
}
