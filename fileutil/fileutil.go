package fileutil

import (
	"errors"
	"io"
	"os"
	fpath "path/filepath"
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
  - The number of bytes that were copied (zero in case there was an error)
  - An error if one occurred or nil if no error occurred.
*/
func CopyFile(sourceFile string, destinationFileOrDirectory string, overwrite bool) (createdFile string, fileSize int64, err error) {
	// stat the source file
	stat, err := os.Stat(sourceFile)
	if err != nil {
		return "", 0, errors.New("Failed to stat source file (" + sourceFile + "): " + err.Error())
	}

	// check if it is a regular file and not a directory
	if !stat.Mode().IsRegular() {
		return "", 0, errors.New(sourceFile + " is not a regular file")
	}

	destinationPath := destinationFileOrDirectory

	// check if the destination file is a directory and if so, append the source file name
	stat, err = os.Stat(destinationPath)
	if err != nil && !os.IsNotExist(err) {
		// we got an error and the error was not due to the path not existing
		return "", 0, errors.New("Failed to stat destination path (" +
			destinationPath + "): " + err.Error())

	} else if err != nil {
		// we have a non-existent file path -- check that its parent directory exists
		dir := fpath.Dir(destinationPath)
		dirStat, err := os.Stat(dir)
		if err != nil {
			return "", 0, errors.New("Failed to stat directory (" + dir +
				") for destination path (" + destinationPath +
				"): " + err.Error())
		}
		if !dirStat.Mode().IsDir() {
			return "", 0, errors.New("Directory (" + dir + ") for destination path (" +
				destinationPath + ") is not a directory.")
		}

	} else if stat.Mode().IsDir() {
		// the destination is a directory so append the file name
		destinationPath = fpath.Join(destinationPath, fpath.Base(sourceFile))

		// if not allowing overwrite, check if the new path exists
		if !overwrite {
			// check if the new file path exists
			_, err = os.Stat(destinationPath)
			if err == nil || (err != nil && !os.IsNotExist(err)) {
				return "", 0, errors.New("The target file already exists in the destination " +
					"directory and overwrite is not allowed: " + destinationPath)
			}
		}

	} else if !overwrite {
		return "", 0, errors.New("Destination file already exists and overwrite is not allowed: " +
			destinationPath)
	}

	// open the source file
	source, err := os.Open(sourceFile)
	if err != nil {
		return "", 0, errors.New("Failed to open source file (" + sourceFile + "): " + err.Error())
	}
	defer source.Close() // defer closing the source file

	// create the destination file
	destination, err := os.Create(destinationPath)
	if err != nil {
		return "", 0, errors.New("Failed to create destination file (" + destinationPath +
			"): " + err.Error())
	}
	defer destination.Close()

	// copy the data from source to destination
	byteCount, err := io.Copy(destination, source)
	return destinationPath, byteCount, err
}