# Entityboard-Go-Adapter

A Go package to help create entityboard-adapters.

Manifest Template

```json
{
    "ID": "company.plugin",   // Unique id of the plugin
    "Name": "Example Plugin", // Display name

    // Item Types provided by this plugin
    "Types": [
        {
            "ID": "myItem",         // Id of type. Must be unique per plugin
            "Name": "My Test Item", // Display name
            "Display": "Default",   // Display position: see 'Display' section below
            // Actions supported by this type
            "Actions": [
                "CREATE",
                "UPDATE",
                "DELETE",
                "UPLOAD",
            ],
            // Template of the Datastructure
            // All top-level properties must be defined here
            "Data": {
                "FilePath": "string"
            },
            // Template Items that can be created from the UI (only makes sense with CREATE action)
            "Templates": [
                {
                    "Name": "Template 1",
                    "Data": {
                        "FilePath": "/home/example/example.txt"
                    }
                }
            ]
        }
    ]
}
```

## Display Options
Default: Item is displayed as a regular icon

Options: Item is displayed in the options menu

Hidden: Item is never displayed in UI

## Commands

Commands are handled per item type or in default handlers.

### Available Commands

- get-items: []int: return all items for the ids in the data. If an item is missing, it will be ignored.

- enumerate-items: {limit int, offset int, type string}: returns limit (or less if less available) amount of items, in application defined order, offset by offset. Only returns items of the specified type.

- activate: int: activates the item with the specified id. Doesn't return anything except for errors.

- create: {TypeId string, Name string, Data object}: creates a new item of the specified type, with the specified name and data. The adapter then returns the id of the new item. This command only needs to be implemented, when the CREATE action is defined on the type.

- update: {id int, Data object}: updates the data of the item with the specified id. This command only needs to be implemented, when the UPDATE action is defined on the type.

- delete: []int: deletes all items whose id is in the passed array. This command only needs to be implemented, when the DELETE action is defined on the type.

- upload: {id int, data BSON, fileName string}: uploads a file to an item. Each upload is associated with an item, specified by id. The file data is passed as BSON, the file name as a string. This command only needs to be implemented, when the UPLOAD action is defined on the type.