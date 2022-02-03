# Entityboard-Go-Adapter

A Go package to help create entityboard-adapters.

Manifest Template

```js
{
    "ID": "company.plugin",   // Unique id of the plugin
    "Name": "Example Plugin", // Display name
    "Socket": "example_socket",

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

A Command looks like this (JSON)

```js
// Example: activate
// This command would activate the item with id 1.
{
    "Verb": "activate",
    "ItemType": "testitem",
    "Noun": null, // optional
    "ID": 1       // optional
}

// Example: enumerate-items
// This command would return the first 10 items of type testitem.
{
    "Verb": "enumerate-items",
    "ItemType": "testitem",
    "Noun": {
        "Offset": 0,
        "Limit": 10
    },
    "ID": null
}
```

### Notes:

The 'ItemType' field technically redundant for commands that pass an id,
Since the adapter could determine the type by itself. It is still passed by
the server since it usually knows the item before calling a command on it.

Item ids have to be unique, even if the items have different types.

### Available Commands

- get-item: Returns the item. If the item doesn't exist, this returns null.

- enumerate-items: {limit int, offset int}: Returns limit (or less if less available) amount of items, in application defined order, offset by offset. 
(Does not use Id field)

- activate: int: Activates the item with the specified id. Doesn't return anything except for errors.

- create: {Name string, Data object}: Creates a new item of the specified type, with the specified name and data. The adapter then returns the id of the new item. This command only needs to be implemented, when the CREATE action is defined on the type.
(Does not use Id field)

- update: {Name string, Data object}: Updates the data of the item. This command only needs to be implemented, when the UPDATE action is defined on the type.

- delete: Deletes the item. This command only needs to be implemented, when the DELETE action is defined on the type.

- upload: {file BSON, fileName string}: Uploads a file to an item. Each upload is associated with an item, specified by id. The file data is passed as BSON, the file name as a string. This command only needs to be implemented, when the UPLOAD action is defined on the type.
