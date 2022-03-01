package adapter

// EnumerateItemsData is the noun for enumerate-items
type EnumerateItemsData struct {
	Offset uint64
	Limit  uint64
}

// Item represents an instance of an item type.
// These are the items that get displayed in the UI.
type Item struct {
	ID       uint64
	Name     string
	ItemType string
	Icon     string
	Color    string

	Data interface{}
}

// Enum should be returned in a slice as a response to get-enum
type Enum struct {
	Text  string
	Value int
}
