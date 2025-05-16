package tls

import (
	"crypto/tls"

	"github.com/senzing-garage/go-helpers/wraperror"
)

/*
Return an TLS Certificate from a certificte file and an encrypted or unencrypted key file.

Input
  - certFile: The path to the regular file that is the source for the copying.
  - keyFile: The key file for the certificate.
  - password: The password for the keyFile.  If empty string, the keyFile is assumed to be unencrypted.

Output
  - A TLS certificate
*/
func LoadX509KeyPair(certFile string, keyFile string, password string) (tls.Certificate, error) {
	if len(password) == 0 {
		result, err := tls.LoadX509KeyPair(certFile, keyFile)

		return result, wraperror.Errorf(err, "length=0")
	}

	result, err := loadX509KeyPairWithPassword(certFile, keyFile, password) // OS-specific implementation.

	return result, wraperror.Error(err)
}
