package adapter

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
)

type Command struct {
	// Id is the ItemID of the item, on which this command is called.
	// If the command is not called on an item (e.g. enumerate-items), this is nil.
	Id *int

	// ItemType is the ItemType of the targeted items.
	// I.e. calling enumerate-items will only return items of the type specified here.
	ItemType string

	// Verb is the requested action (i.e enumerate-items or activate)
	Verb string

	// Noun is additional data passed for some verbs.
	// E.g { Offset int, Limit int} for enumerate-items or nil for activate.
	Noun interface{}
}

func WriteError(w io.Writer, err error) {
	str := err.Error()
	errJson, err := json.Marshal(struct {
		Err string
	}{str})

	if err != nil {
		log.Println(err)
		return
	}

	if _, err := fmt.Fprintf(w, "%s\000", string(errJson)); err != nil {
		log.Println(err)
	}
}

func WriteJson(w io.Writer, data interface{}) error {
	dataJson, err := json.Marshal(data)

	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(w, "%s\000", string(dataJson))

	return err
}
