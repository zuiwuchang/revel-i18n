package cmdnew

// Item 每個待 翻譯條目
type Item struct {
	Key   string
	Value string

	Description map[string]bool
}

// Items .
type Items []*Item

// Len is the number of elements in the collection.
func (items Items) Len() int {
	return len(items)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (items Items) Less(i, j int) bool {
	return items[i].Key < items[j].Key
}

// Swap swaps the elements with indexes i and j.
func (items Items) Swap(i, j int) {
	items[i], items[j] = items[j], items[i]
}
