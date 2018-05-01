package hashing

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
)

type hashTableTestCases struct {
	tag            string
	inputValues    []([]int)
	expectedValues []int
}

func TestHashFunction(t *testing.T) {
	fmt.Println("Running TestHashFunction...")
	capacity := 20
	hashTable := NewHashTable(capacity)
	// hashTable.HashFunction = MultiplicativeHash
	for i := 0; i < 40; i++ {
		expectedHashValue := i %  capacity
		actualHashValue := hashTable.HashFunction(i, capacity)
		if actualHashValue != expectedHashValue {
			t.Errorf(
				"Error hashing value %d. Expect hash value %d, received hash value %d",
				i, expectedHashValue, actualHashValue,
			)
		}
	}
}

func TestHashTable(t *testing.T) {
	fmt.Println("Running TestHashTable...")
	testCases := []hashTableTestCases{
		{
			tag: "Case 1",
			inputValues: [][]int{
				[]int{1, 2},
				[]int{4, 2},
			},
			expectedValues: []int{},
		},
	}

	hashTable.Set(NewHashableInt(rand.Int()))

	for _, testCase := range testCases {
		if !reflect.DeepEqual(testCase, testCase) {
			t.Errorf("PLACEHODLER test failed")
		}
	}
}
