package adapter

// EnumerateItemsData is the noun for enumerate-items
type EnumerateItemsData struct {
	Offset int
	Limit  int
}

// CreateData is the noun for create
type CreateData struct {
	Name string
	Data interface{}
}

// UpdateData is the noun for update
type UpdateData struct {
	Name string
	Data interface{}
}

// UploadData is the noun for upload
type UploadData struct {
	File     []byte
	FileName string
}
