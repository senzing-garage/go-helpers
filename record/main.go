package record

import (
	"github.com/senzing-garage/go-messaging/messenger"
)

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

// Identifier of this package for the error string "SZSDK6403xxxx".
const ComponentID = 6403

// ----------------------------------------------------------------------------
// Variables
// ----------------------------------------------------------------------------

// Message templates.
var IDMessages = map[int]string{
	3000: "JSON-line not well formed",
	3001: "a DATA_SOURCE field is required",
	3002: "a RECORD_ID field is required",
}

var messengerOptions = []interface{}{}
var szerrors, _ = messenger.New(messengerOptions...)
