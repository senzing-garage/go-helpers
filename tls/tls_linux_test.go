//go:build linux

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
	test.Parallel()

	certFile := "../testdata/certificates/client/certificate.pem"
	keyFile := "../testdata/certificates/client/private_key_encrypted.pem"
	password := "Passw0rd"
	_, err := localtls.LoadX509KeyPair(certFile, keyFile, password)
	require.NoError(test, err)
}
