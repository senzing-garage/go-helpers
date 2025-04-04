package tls_test

import (
	"testing"

	localtls "github.com/senzing-garage/go-helpers/tls"
	"github.com/stretchr/testify/require"
)

// ----------------------------------------------------------------------------
// Test Public functions
// ----------------------------------------------------------------------------

func TestLoadX509KeyPair_unencrypted(test *testing.T) {
	certFile := "../testdata/certificates/client/certificate.pem"
	keyFile := "../testdata/certificates/client/private_key.pem"
	password := ""
	_, err := localtls.LoadX509KeyPair(certFile, keyFile, password)
	require.NoError(test, err)
}
