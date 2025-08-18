//go:build windows

package tls

import (
	"context"
	"crypto/tls"
	"fmt"

	"github.com/senzing-garage/go-helpers/wraperror"
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
	return tlsCertificate, wraperror.Errorf(
		errForPackage,
		fmt.Sprintf("cannot decript %s on Windows platform", keyFile),
	)
}
