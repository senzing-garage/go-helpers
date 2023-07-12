//go:build darwin

package g2engineconfigurationjson

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// ----------------------------------------------------------------------------
// Internal functions
// ----------------------------------------------------------------------------

func testError(test *testing.T, err error) {
	if err != nil {
		assert.FailNow(test, err.Error())
	}
}

// ----------------------------------------------------------------------------
// Test interface functions
// ----------------------------------------------------------------------------

func TestBuildSimpleSystemConfigurationJson(test *testing.T) {
	_, err := BuildSimpleSystemConfigurationJson("postgresql://postgres:postgres@$10.0.0.1:5432/G2")
	testError(test, err)
}
