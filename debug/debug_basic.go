/*
Package debug is used to assist in debugging code.
*/
package debug

import "fmt"

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// BasicDebug is the default implementation of the Debug interface.
type BasicDebug struct {
	MinimumDebugLevel int
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

/*
The Print wraps fmt.Print with guard that is based on BasicDebug.DebugLevel.

Input
  - level:  Debug level. The higher the number, the less likely it will be printed
  - a: A series of variables to print.
*/
func (debug *BasicDebug) Print(level int, a ...any) {
	if debug.MinimumDebugLevel > 0 {
		if level >= debug.MinimumDebugLevel {
			fmt.Print(a...) //nolint
		}
	}
}

/*
The Printf wraps fmt.Printf with guard that is based on BasicDebug.DebugLevel.

Input
  - level:  Debug level. The higher the number, the less likely it will be printed
  - format: A string format specifier.
  - a: A series of variables to insert into the format
*/
func (debug *BasicDebug) Printf(level int, format string, a ...any) {
	if debug.MinimumDebugLevel > 0 {
		if level >= debug.MinimumDebugLevel {
			fmt.Printf(format, a...) //nolint
		}
	}
}

/*
The Println wraps fmt.Println with guard that is based on BasicDebug.DebugLevel.

Input
  - level:  Debug level. The higher the number, the less likely it will be printed
  - a: A series of variables to print.
*/
func (debug *BasicDebug) Println(level int, a ...any) {
	if debug.MinimumDebugLevel > 0 {
		if level >= debug.MinimumDebugLevel {
			fmt.Println(a...) //nolint
		}
	}
}
