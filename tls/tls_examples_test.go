//go:build linux

package tls_test

import (
	"fmt"

	localtls "github.com/senzing-garage/go-helpers/tls"
)

// ----------------------------------------------------------------------------
// Examples for godoc documentation
// ----------------------------------------------------------------------------

func ExampleLoadX509KeyPair() {
	const (
		certFile = "../testdata/certificates/client/certificate.pem"
		keyFile  = "../testdata/certificates/client/private_key.pem"
	)

	var password string

	_, err := localtls.LoadX509KeyPair(certFile, keyFile, password)
	if err != nil {
		fmt.Println(err)
	}
	// Output:
}
