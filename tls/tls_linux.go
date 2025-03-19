//go:build linux

package tls

import (
	"crypto/tls"
	"fmt"
	"os"
	"os/exec"
)

// ----------------------------------------------------------------------------
// OS specific functions.
// ----------------------------------------------------------------------------

func loadX509KeyPairWithPassword(certFile string, keyFile string, password string) (tls.Certificate, error) {
	var err error
	var tlsCertificate tls.Certificate

	tmpFile, err := os.CreateTemp("", "tmp_tls_decrypted_key_")
	if err != nil {
		return tlsCertificate, err
	}
	defer os.Remove(tmpFile.Name())

	path, err := exec.LookPath("openssl")
	if err != nil {
		return tlsCertificate, err
	}
	passin := fmt.Sprintf("pass:%s", password)
	cmd := exec.Command(path, "rsa", "-in", keyFile, "-out", tmpFile.Name(), "-passin", passin)
	_, err = cmd.Output()
	if err != nil {
		return tlsCertificate, err
	}
	return tls.LoadX509KeyPair(certFile, tmpFile.Name())
}
