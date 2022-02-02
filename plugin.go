package adapter

import (
	"log"

	"github.com/NicoKleinschmidt/entityboard-go-adapter/pipe"
)

type Plugin struct {
	// ID of the plugin
	ID string

	// Name of the plugin
	Name string

	// NamedPipe is the path to the named pipe.
	NamedPipe string

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
// It starts listening for and accepting connections to the named pipe.
// This function will not return unless an error occurs.
func (pl *Plugin) Start() error {
	pipe := pipe.NamedPipe{
		Path: pl.NamedPipe,
	}

	if err := pipe.Open(); err != nil {
		return err
	}

	for {
		conn, err := pipe.WaitForConnection()

		if err != nil {
			log.Println(err)
			continue
		}

		go pl.commandHandler(conn)
	}
}
