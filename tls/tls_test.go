package tls

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// ----------------------------------------------------------------------------
// Test Public functions
// ----------------------------------------------------------------------------

func TestLoadX509KeyPair_encrypted(test *testing.T) {
	certFile := "../testdata/certificates/client/certificate.pem"
	keyFile := "../testdata/certificates/client/private_key_encrypted.pem"
	password := "Passw0rd"
	_, err := LoadX509KeyPair(certFile, keyFile, password)
	require.NoError(test, err)
}
