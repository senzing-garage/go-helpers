//go:build linux

package tls

import (
	"crypto/tls"
	"os"
	"os/exec"

	"github.com/senzing-garage/go-helpers/wraperror"
)

// ----------------------------------------------------------------------------
// OS specific functions.
// ----------------------------------------------------------------------------

func loadX509KeyPairWithPassword(certFile string, keyFile string, password string) (tls.Certificate, error) {
	var err error

	var tlsCertificate tls.Certificate

	tmpFile, err := os.CreateTemp("", "tmp_tls_decrypted_key_")
	if err != nil {
		return tlsCertificate, wraperror.Errorf(err, "tls.loadX509KeyPairWithPassword.os.CreateTemp error: %w", err)
	}
	defer os.Remove(tmpFile.Name())

	path, err := exec.LookPath("openssl")
	if err != nil {
		return tlsCertificate, wraperror.Errorf(err, "tls.loadX509KeyPairWithPassword.exec.LookPath error: %w", err)
	}

	passin := "pass:" + password
	cmd := exec.Command(path, "rsa", "-in", keyFile, "-out", tmpFile.Name(), "-passin", passin)

	_, err = cmd.Output()
	if err != nil {
		return tlsCertificate, wraperror.Errorf(err, "tls.loadX509KeyPairWithPassword.cmd.Output error: %w", err)
	}

	result, err := tls.LoadX509KeyPair(certFile, tmpFile.Name())

	return result, wraperror.Errorf(err, "tls.loadX509KeyPairWithPassword error: %w", err)
}
