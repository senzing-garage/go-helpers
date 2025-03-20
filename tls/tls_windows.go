//go:build windows

package tls

import (
	"crypto/tls"
	"fmt"
)

func loadX509KeyPairWithPassword(certFile string, keyFile string, password string) (tls.Certificate, error) {
	_ = certFile
	_ = keyFile
	_ = password
	var tlsCertificate tls.Certificate
	return tlsCertificate, fmt.Errorf("cannot decript %s on Windows platform", keyFile)
}
