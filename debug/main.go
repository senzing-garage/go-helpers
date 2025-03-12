// The debug package helps with debugging.
package debug

// ----------------------------------------------------------------------------
// Types - interface
// ----------------------------------------------------------------------------

type Debug interface {
	Print(level int, a ...any)
	Printf(level int, format string, a ...any)
	Println(level int, a ...any)
}

// ----------------------------------------------------------------------------
// Constructor methods
// ----------------------------------------------------------------------------

func New(level int) (Debug, error) {
	var err error
	result := &BasicDebug{
		MinimumDebugLevel: level,
	}
	return result, err
}
