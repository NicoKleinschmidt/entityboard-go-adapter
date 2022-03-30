package adapter

// EnumerateItemsData is the noun for enumerate-items
type EnumerateItemsData struct {
	Offset     uint64
	Limit      uint64
	SearchText string
	Tags       []string
}

// GetImageNoun is the noun for get-items
type GetImageNoun struct {
	Path string
}
