package adapter

import (
	"log"

	entityipc "github.com/NicoKleinschmidt/entity-ipc"
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
// It starts listening for and accepting connections to the socket.
// This function will not return unless an error occurs.
func (pl *Plugin) Start() error {
	pl.defaultHandler = handlerInfo{
		handlerFunc: defaultHandler,
	}

	ls, err := entityipc.Listen(pl.Socket)

	if err != nil {
		return err
	}

	for {
		conn, err := ls.AcceptJson()

		if err != nil {
			log.Println(err)
			continue
		}

		go pl.commandHandler(conn)
	}
}
