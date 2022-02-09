package adapter

import (
	"encoding/json"
	"fmt"
)

type PluginManifest struct {
	ID     string
	Name   string
	Socket string

	Settings        interface{}
	DefaultSettings interface{}

	Types []ItemType
}

// ParseManifest parses a plugin manifest and creates a plugin of the configuration.
func ParseManifest(manifest PluginManifest) (*Plugin, error) {
	if manifest.ID == "" {
		return nil, fmt.Errorf("plugin id must not be empty")
	}

	plugin := Plugin{
		ID:     manifest.ID,
		Name:   manifest.Name,
		Socket: manifest.Socket,
	}

	plugin.ItemTypes = map[string]*ItemType{}

	for _, itemType := range manifest.Types {
		itemTypeCopy := itemType
		plugin.ItemTypes[itemType.ID] = &itemTypeCopy
	}

	return &plugin, nil
}

// ParseManifestFile does the same as ParseManifest, except that its parameter is the json data.
func ParseManifestFile(data []byte) (*Plugin, error) {
	var manifest PluginManifest

	if err := json.Unmarshal(data, &manifest); err != nil {
		return nil, err
	}

	return ParseManifest(manifest)
}
