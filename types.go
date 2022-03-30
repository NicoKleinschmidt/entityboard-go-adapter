package adapter

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

// ImagesManifest is the 'Images' object of the plugin manifest.
type ImagesManifest struct {
	Static  ImagesStaticManifest
	Dynamic ImagesDynamicManifest
}

// ImagesStaticManifest is the 'Static' object of the 'Images' object.
type ImagesStaticManifest struct {
	Enabled bool
	Path    string
}

// ImagesDynamicManifest is the 'Dynamic' object of the 'Images' object.
type ImagesDynamicManifest struct {
	Enabled bool
	Prefix  string
}
