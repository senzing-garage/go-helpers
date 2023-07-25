package jsonutil

import (
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

func TestNormalizeJson_Basic(test *testing.T) {
	var jsonText = "{\"foo\": 123, \"bar\": \"abc\", \"phoo\": true, \"lum\": 20.5}"
	var expected = "{\"bar\":\"abc\",\"foo\":123,\"lum\":20.5,\"phoo\":true}"
	actual, err := NormalizeJson(jsonText)
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON not normalized as expected")
}
