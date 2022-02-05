package adapter

// EnumerateItemsData is the noun for enumerate-items
type EnumerateItemsData struct {
	Offset int
	Limit  int
}

// Item represents an instance of an item type.
// These are the items that get displayed in the UI.
type Item struct {
	ID   uint64
	Name string

	Data interface{}
}
