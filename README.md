# Entityboard-Go-Adapter

A Go package to help create entityboard-adapters.

Manifest Template

```js
{
    "ID": "company.plugin",   // Unique id of the plugin
    "Name": "Example Plugin", // Display name
    "Socket": "example_socket",

    // Settings template
    // Parsed the same as itemType data template
    "Settings": [

    ],

    // Item Types provided by this plugin
    "Types": [
        {
            "ID": "myItem",         // Id of type. Must be unique per plugin
            "Name": "My Test Item", // Display name
            "Component": "button",

            // Actions supported by this type
            "Actions": [
                "CREATE",
                "UPDATE",
                "DELETE",
            ],
            // Template of the Datastructure
            // Parsed the same as 'Settings'
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

## Settings template format

```json
[
    {
        "field": "exampleString",
        "label": "Example String",
        "icon": "",
        "desc": "Example String Description",
        "type": "string"
    },
    {
        "label": "More Settings",
        "icon": "",
        "desc": "Submenu",
        "type": [
            {
                "field": "exampleEnum",
                "label": "Example Enum",
                "icon": "",
                "desc": "Example Enum Description",
                "type": "emum:example"
            },
            {
                "field": "submenu.exampleStringArr",
                "label": "Example String Array",
                "icon": "",
                "desc": "Example String Array Description",
                "type": "string[]"
            }
        ]
    }
]

```

Possible types: string, number, boolean, object, file, enum + array versions (i.e. string[], enum[]:id).

- string: accepts string values

- number: accepts number values

- boolean: accepts boolean values

- object: key-value pairs

- file: accepts files in this format:

```json
    {
        "data": "base64 string",
        "filename": "filename"
    }
```

- arrays: allows multiple values of the defined type

- enum: enums end with an identifier. (enum:identifier, enum[]:identifier)
This identifier is used to query the possible values. This is done by sending a 'get-enum' command to the plugin.
The command is never associated with an item type (even for the item data template). The noun of this command is a string containing the identifier.
The plugin should then respond with an array objects like this:
```json
[
    {
        "Text": "Option 1",
        "Value": 0
    },
    {
        "Text": "Option 2",
        "Value": 1
    }
]
```
When the settings get applied, the objects value for this key will be the value field of the selected option (only the number for single enums, number array for enum arrays). Enum arrays can be used like flags.

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

- apply-settings: {Settings}: Applies the settings passed as the noun. The plugin can respond in 3 ways: no response -> nothing happens, "reload" -> the server will reload all items or "restart" -> the server will stop and restart the plugin.
(Does not use ItemType or Id fields)

- get-settings: {Settings}: Get the current settings.
(Does not use ItemType or Id fields)

- get-enum: identifier: Get the possible values for the specified enum identifier.
(Does not use ItemType or Id fields)