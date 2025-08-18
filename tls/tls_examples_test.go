//go:build linux

package tls_test

import (
	"context"
	"fmt"

	"github.com/senzing-garage/go-helpers/tls"
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

	ctx := context.Background()

	tlsCertificate, err := tls.LoadX509KeyPair(ctx, certFile, keyFile, password)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(len(tlsCertificate.Certificate))
	// Output: 1
}
