package adapter

// EnumerateItemsData is the noun for enumerate-items
type EnumerateItemsData struct {
	Offset     uint64
	Limit      uint64
	SearchText string
	Tags       []string
}

// Item represents an instance of an item type.
// These are the items that get displayed in the UI.
type Item struct {
	ID       uint64
	Name     string
	ItemType string
	Icon     string
	Color    string

	// Tags for this item.
	// These should be tags defined for this item type
	// and the same tags, this item gets filtered by.
	Tags []string

	Data interface{}
}

// Enum should be returned in a slice as a response to get-enum
type Enum struct {
	Text  string
	Value int
}
