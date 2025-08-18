//go:build windows

package tls

import (
	"context"
	"crypto/tls"
	"fmt"
)

func loadX509KeyPairWithPassword(
	ctx context.Context,
	certFile string,
	keyFile string,
	password string,
) (tls.Certificate, error) {
	_ = ctx
	_ = certFile
	_ = keyFile
	_ = password
	var tlsCertificate tls.Certificate
	return tlsCertificate, fmt.Errorf("cannot decript %s on Windows platform", keyFile)
}
