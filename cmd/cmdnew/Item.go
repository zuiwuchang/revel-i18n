package cmdnew

// Item 每個待 翻譯條目
type Item struct {
	Key   string
	Value string

	Description map[string]bool
}
