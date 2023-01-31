package observer

import (
	"context"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

type Observer interface {
	Update(ctx context.Context, message string)
}
