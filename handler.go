package adapter

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"reflect"
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
	reader := bufio.NewReader(conn)
	defer conn.Close()

	for {
		msg, err := reader.ReadBytes('\000')

		if err != nil {
			log.Println(err)
			return
		}

		var cmd commandRawData

		if err := json.Unmarshal(msg[:len(msg)-1], &cmd); err != nil {
			WriteError(conn, err)
			continue
		}

		response, err := pl.findAndCallHandler(cmd)

		if err != nil {
			WriteError(conn, err)
			continue
		}

		if err := WriteJson(conn, response); err != nil {
			log.Println(err)
		}
	}
}

// findAndCallHandler finds and calls the correct handler function for the passed command.
// returns an error if the itemType specified doesn't exist.
func (pl Plugin) findAndCallHandler(cmd commandRawData) (interface{}, error) {
	if itemType, ok := pl.ItemTypes[cmd.ItemType]; ok {
		if handler, ok := itemType.handlers[cmd.Verb]; ok {
			return handler.call(cmd)
		}
	} else {
		return nil, fmt.Errorf("404: ItemType not found")
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
