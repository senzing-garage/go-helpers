package observed

import (
	"context"

	"github.com/senzing/go-interfaces/observer"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

type Observed interface {
	RegisterObserver(ctx context.Context, observer observer.Observer)
	UnregisterObserver(ctx context.Context, observer observer.Observer)
}
