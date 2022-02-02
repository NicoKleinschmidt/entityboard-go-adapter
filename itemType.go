package adapter

type ItemType struct {
	handlers map[string]handlerInfo
}

// Handler registers a command handler for a specific verb on this type.
// data is the type of data that gets passed for this. Can be nil.
// The handler will be called when a command with the specified verb gets called
// on this item type.
//
// The data field of the passed Command struct can be safely asserted to the
// type of the passed data interface to this function.
func (t *ItemType) Handle(verb string, data interface{}, fn CommandHandlerFunc) {
	if t.handlers == nil {
		t.handlers = make(map[string]handlerInfo)
	}

	t.handlers[verb] = handlerInfo{
		data:        data,
		handlerFunc: fn,
	}
}
