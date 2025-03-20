package tls

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// ----------------------------------------------------------------------------
// Test Public functions
// ----------------------------------------------------------------------------

func TestLoadX509KeyPair_unencrypted(test *testing.T) {
	certFile := "../testdata/certificates/client/certificate.pem"
	keyFile := "../testdata/certificates/client/private_key.pem"
	password := ""
	_, err := LoadX509KeyPair(certFile, keyFile, password)
	require.NoError(test, err)
}
