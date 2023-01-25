package record

import (
	"github.com/senzing/go-logging/messagelogger"
)

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

// Identifier of this package for the error string "senzing-6040xxxx".
const ProductId = 6040

// ----------------------------------------------------------------------------
// Variables
// ----------------------------------------------------------------------------

// Message templates.
var IdMessages = map[int]string{
	3000: "JSON-line not well formed",
	3001: "a DATA_SOURCE field is required",
	3002: "a RECORD_ID field is required",
}

var szerrors, _ = messagelogger.NewSenzingApiLogger(ProductId, IdMessages, nil, 0)
