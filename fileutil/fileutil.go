package fileutil

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/senzing-garage/go-helpers/wraperror"
)

/*
Copies a file from the specified source path to the specified destination path which may
be the full destination file path or the path to the directory in which the destination
file should be created with the same simple file name as the source file.

Input
  - sourceFile: The path to the regular file that is the source for the copying.
  - destinationFileOrDirectory: The path to the destination file or the directory in
    which the destination file should be created with the same name as the source file.
  - overwrite: Set to true if you want to allow overwriting an existing destination
    file or false if you want to prevent overwriting an existing destination file.

Output
  - The path to the file that was created or ovewritten (more useful when copying to a directory)
  - The number of bytes that were copied (zero in case there was an error)
  - An error if one occurred or nil if no error occurred.
*/
func CopyFile(sourceFile string, destinationFileOrDirectory string, overwrite bool) (string, int64, error) {
	// Stat the source file.
	stat, err := os.Stat(sourceFile)
	if err != nil {
		return "", 0, fmt.Errorf("failed to stat source file (%v): %v", sourceFile, err.Error())
	}

	// Check if it is a regular file and not a directory.

	if !stat.Mode().IsRegular() {
		return "", 0, fmt.Errorf("%v is not a regular file", sourceFile)
	}

	destinationPath := destinationFileOrDirectory

	// check if the destination file is a directory and if so, append the source file name
	stat, err = os.Stat(destinationPath)

	switch {
	case err != nil && !errors.Is(err, fs.ErrNotExist):
		// we got an error and the error was not due to the path not existing
		return "", 0, fmt.Errorf("failed to stat destination path (%v): %v", destinationPath, err.Error())
	case err != nil:
		// we have a non-existent file path -- check that its parent directory exists
		dir := filepath.Dir(destinationPath)
		dirStat, err := os.Stat(dir)

		if err != nil {
			return "", 0, fmt.Errorf(
				"failed to stat directory (%s) for destination path (%s): %s",
				dir,
				destinationPath,
				err.Error(),
			)
		}

		if !dirStat.Mode().IsDir() {
			return "", 0, fmt.Errorf(
				"directory (%v) for destination path (%v) is not a directory",
				dir,
				destinationPath,
			)
		}
	case stat.Mode().IsDir():
		// the destination is a directory so append the file name
		destinationPath = filepath.Join(destinationPath, filepath.Base(sourceFile))

		// if not allowing overwrite, check if the new path exists
		if !overwrite {
			// check if the new file path exists
			_, err = os.Stat(destinationPath)
			if err == nil || (err != nil && !errors.Is(err, fs.ErrNotExist)) {
				return "", 0, fmt.Errorf("the target file already exists in the destination "+
					"directory and overwrite is not allowed: %v", destinationPath)
			}
		}
	case !overwrite:
		return "", 0, fmt.Errorf("destination file already exists and overwrite is not allowed: %v", destinationPath)
	}

	byteCount, err := copyFile(sourceFile, destinationPath)

	return destinationPath, byteCount, wraperror.Errorf(err, "fileutil.CopyFile error: %w", err)
}

// ----------------------------------------------------------------------------
// Private functions
// ----------------------------------------------------------------------------

func copyFile(sourceFile string, destinationPath string) (int64, error) {
	// open the source file
	source, err := os.Open(filepath.Clean(sourceFile))
	if err != nil {
		return 0, fmt.Errorf("failed to open source file ("+sourceFile+"): %v", err.Error())
	}
	defer source.Close() // defer closing the source file

	// create the destination file
	destination, err := os.Create(filepath.Clean(destinationPath))
	if err != nil {
		return 0, fmt.Errorf("failed to create destination file (%v): %v", destinationPath, err.Error())
	}
	defer destination.Close()

	// copy the data from source to destination
	byteCount, err := io.Copy(destination, source)

	return byteCount, wraperror.Errorf(err, "io.Copy error: %w", err)
}
