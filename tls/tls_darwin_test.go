//go:build darwin

package tls_test

import (
	"testing"

	localtls "github.com/senzing-garage/go-helpers/tls"
	"github.com/stretchr/testify/require"
)

// ----------------------------------------------------------------------------
// Test Public functions
// ----------------------------------------------------------------------------

func TestLoadX509KeyPair_encrypted(test *testing.T) {

	ctx := test.Context()

	certFile := "../testdata/certificates/client/certificate.pem"
	keyFile := "../testdata/certificates/client/private_key_encrypted.pem"
	password := "Passw0rd"
	_, err := localtls.LoadX509KeyPair(ctx, certFile, keyFile, password)
	require.Error(test, err)
}
