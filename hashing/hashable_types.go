package hashing

type HashableInt int

func NewHashableInt(value int) HashableInt {
	return HashableInt(value)
}

func (h HashableInt) Key() int {
	return int(h)
}

func (h HashableInt) Value() int {
	return int(h)
}
