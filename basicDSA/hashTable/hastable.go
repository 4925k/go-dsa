package hashtable

type hashTable map[int]any

func hashFunction(key int) int {
	// use proper function
	// for now we just do a basic multiplication
	return key * 17
}

func (h *hashTable) Insert(key int, data any) {
	idx := hashFunction(key)
	(*h)[idx] = data
}

func (h *hashTable) Remove(key int) {
	(*h)[key] = nil
}

func hashtableTest() {
	var ht hashTable

	ht.Insert(123, "hello world")

	ht.Remove(123)
}
