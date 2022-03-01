package adapter

import (
	"encoding/json"
	"fmt"
	"net"
	"reflect"

	ipc "github.com/NicoKleinschmidt/entity-ipc"
)

type handlerInfo struct {
	data        interface{}
	handlerFunc CommandHandlerFunc
}

type CommandHandlerFunc func(cmd Command) (interface{}, error)

// commandRawData is the Command struct, except is has an additional
// Data field, that is still in the JSON format.
//
// This is used to later unmarshal the data into the correct type.
type commandRawData struct {
	Command
	RawData []byte
}

func (cmd *commandRawData) UnmarshalJSON(data []byte) (err error) {
	if err := json.Unmarshal(data, &cmd.Command); err != nil {
		return err
	}

	cmd.RawData, err = json.Marshal(cmd.Noun)

	return
}

// commandHandler is the handler for incomming connections from the named pipe.
// This should be called as a goroutine for all accepted connections.
//
// This function will not return until the server ends the connection.
func (pl Plugin) commandHandler(conn net.Conn) {
	defer conn.Close()

	ipcHandler := ipc.IPC{}
	ipcHandler.Start(conn)

	ipcHandler.Handle(commandRawData{}, func(cmd interface{}) (interface{}, error) {
		response, err := pl.findAndCallHandler(cmd.(commandRawData))

		if err != nil {
			return nil, err
		}

		return response, nil
	})
}

// findAndCallHandler finds and calls the correct handler function for the passed command.
// returns an error if the itemType specified doesn't exist.
func (pl Plugin) findAndCallHandler(cmd commandRawData) (interface{}, error) {
	if itemType, ok := pl.ItemTypes[cmd.ItemType]; ok {
		if handler, ok := itemType.handlers[cmd.Verb]; ok {
			return handler.call(cmd)
		}
	}

	if handler, ok := pl.handlers[cmd.Verb]; ok {
		return handler.call(cmd)
	}

	return pl.defaultHandler.call(cmd)
}

// call calls the handler with the specified command.
func (h handlerInfo) call(raw commandRawData) (interface{}, error) {
	cmd := raw.Command

	if h.data != nil {
		dataDst := reflect.New(reflect.TypeOf(h.data)).Elem()
		err := json.Unmarshal(raw.RawData, dataDst.Addr().Interface())

		if err != nil {
			return nil, err
		}

		cmd.Noun = dataDst.Interface()
	}

	return h.handlerFunc(cmd)
}

func defaultHandler(cmd Command) (interface{}, error) {
	return nil, fmt.Errorf("handler not found")
}
