package hashing

import "fmt"

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

func (h *HashTable) Hash(value int) int {
	hashedIndex := h.HashFunction(value, h.Capacity)
	return hashedIndex
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

func (h *HashTable) Add(item Hashable) {
	hashedIndex := h.Hash(item.Key())
	itemValue := item.Value()
	bucket := h.Buckets[hashedIndex]

	if bucket == nil {
		bucket = NewNode(itemValue)
	} else {
		bucket.Add(itemValue)
	}
	h.Buckets[hashedIndex] = bucket
}

func ModuloHash(value, base int) int {
	index := value % base
	return index
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
