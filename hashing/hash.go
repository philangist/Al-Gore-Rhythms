package hashing

import (
	"fmt"
	"math"
)

type Hashable interface {
	Key() int
	Value() int
}

type HashTable struct {
	Capacity     int
	Buckets      []*Node
	HashFunction func(value, base int) int
}

func NewHashTable(capacity int) *HashTable {
	buckets := make([]*Node, capacity)
	return &HashTable{capacity, buckets, ModuloHash}
}

func (h *HashTable) Hash(item Hashable) int {
	hashedIndex := h.HashFunction(item.Key(), h.Capacity)
	return hashedIndex
}

func (h *HashTable) Add(item Hashable) {
	hashedIndex := h.Hash(item)
	itemValue := item.Value()
	bucket := h.Buckets[hashedIndex]

	if bucket == nil {
		bucket = NewNode(itemValue)
	} else {
		bucket.Add(itemValue)
	}
	h.Buckets[hashedIndex] = bucket
}

func (h *HashTable) ItemDistribution() ([]int, int) {
	totalItems := 0
	itemDistribution := make([]int, h.Capacity)
	for index, bucket := range h.Buckets {
		itemDistribution[index] = 0
		if bucket == nil {
			continue
		}
		for {
			totalItems++
			itemDistribution[index]++
			if bucket.Next != nil {
				bucket = bucket.Next
			} else {
				break
			}
		}
	}
	fmt.Println("totalItems is ", totalItems)
	fmt.Println("itemDistribution is ", itemDistribution)
	return itemDistribution, totalItems
}

func (h *HashTable) LoadFactor() float32 {
	_, totalItems := h.ItemDistribution()
	return float32(totalItems) / float32(h.Capacity)
}

func ModuloHash(value, capacity int) int {
	index := value % capacity
	return index
}

func MultiplicativeHash(value, capacity int) int {
	// Based on MIT OCW 6.006 lecture 8 - Hashing with Chaining
	multiplicativeConstant := float64(1954392595985555827)
	rightShiftFactor := int(math.Log2(float64(capacity)))

	value = int(multiplicativeConstant * float64(value))
	value = int(int64(value) % int64(math.Pow(2, 64)))
	value = value >> uint(64 - rightShiftFactor)

	return value
}

type Node struct {
	Value int
	Next  *Node
}

func NewNode(value int) *Node {
	return &Node{value, nil}
}

func (n *Node) Add(value int) {
	for {
		if n.Next == nil {
			n.Next = NewNode(value)
			return
		}
		n = n.Next
	}
}
