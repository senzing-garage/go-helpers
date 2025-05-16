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
		return tlsCertificate, wraperror.Errorf(err, "os.CreateTemp")
	}
	defer os.Remove(tmpFile.Name())

	path, err := exec.LookPath("openssl")
	if err != nil {
		return tlsCertificate, wraperror.Errorf(err, "LookPath")
	}

	passin := "pass:" + password
	cmd := exec.Command(path, "rsa", "-in", keyFile, "-out", tmpFile.Name(), "-passin", passin)

	_, err = cmd.Output()
	if err != nil {
		return tlsCertificate, wraperror.Errorf(err, "cmd.Output")
	}

	result, err := tls.LoadX509KeyPair(certFile, tmpFile.Name())

	return result, wraperror.Error(err)
}
