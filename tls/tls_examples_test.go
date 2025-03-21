//go:build linux

package tls

import "fmt"

// ----------------------------------------------------------------------------
// Examples for godoc documentation
// ----------------------------------------------------------------------------

func ExampleLoadX509KeyPair() {
	var password string
	certFile := "../testdata/certificates/client/certificate.pem"
	keyFile := "../testdata/certificates/client/private_key.pem"
	_, err := LoadX509KeyPair(certFile, keyFile, password)
	if err != nil {
		fmt.Println(err)
	}
	// Output:
}
