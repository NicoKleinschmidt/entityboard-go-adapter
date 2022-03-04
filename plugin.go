package adapter

import (
	"net"
)

type Plugin struct {
	// ID of the plugin
	ID string

	// Name of the plugin
	Name string

	// Socket is the address of the socket this plugin listens on
	Socket string

	// ItemTypeId -> ItemType
	ItemTypes map[string]*ItemType

	// command verb -> command handler
	handlers map[string]handlerInfo

	// defaultHandler will be called if no other handler is found for a command.
	defaultHandler handlerInfo

	conn net.Conn
}

type Command struct {
	// Id is the ItemID of the item, on which this command is called.
	// If the command is not called on an item (e.g. enumerate-items), this is nil.
	Id *uint64

	// ItemType is the ItemType of the targeted items.
	// I.e. calling enumerate-items will only return items of the type specified here.
	ItemType string

	// Verb is the requested action (i.e enumerate-items or activate)
	Verb string

	// Noun is additional data passed for some verbs.
	// E.g { Offset int, Limit int} for enumerate-items or nil for activate.
	Noun interface{}
}

// Handler registers a command handler for a specific verb on this plugin.
// data is the type of data that gets passed for this. Can be nil.
//
// The handler will be called when a command with the specified verb gets called,
// and the ItemType specified in the Command doesn't have a handler registered for the verb.
//
// The data field of the passed Command struct can be safely asserted to the
// type of the passed data interface to this function.
func (pl *Plugin) Handle(verb string, data interface{}, fn CommandHandlerFunc) {
	if pl.handlers == nil {
		pl.handlers = make(map[string]handlerInfo)
	}

	pl.handlers[verb] = handlerInfo{
		data:        data,
		handlerFunc: fn,
	}
}

// Start starts the plugin.
//
// The plugin will try to connect with the server over the socket.
// This function will not until the connection ends.
func (pl *Plugin) Start() error {
	pl.defaultHandler = handlerInfo{
		handlerFunc: defaultHandler,
	}

	conn, err := net.DialUnix("unix", nil, &net.UnixAddr{Name: pl.Socket})

	if err != nil {
		return err
	}

	pl.conn = conn
	pl.commandHandler(conn)

	return nil
}

// Close stops the plugin.
func (pl *Plugin) Close() error {
	if pl.conn == nil {
		return nil
	}
	return pl.conn.Close()
}
