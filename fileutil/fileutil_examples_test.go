package fileutil_test

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/senzing-garage/go-helpers/fileutil"
)

func ExampleCopyFile() {
	// create a file that we will copy (usually this already exists)
	sourceFilePath := filepath.Join(os.TempDir(), "source-file.txt")

	err := os.WriteFile(sourceFilePath, []byte("Hello, World!"), 0600)
	if err != nil {
		fmt.Println(err)
	}

	// define the target path to copy the file to
	targetFilePath := filepath.Join(os.TempDir(), "target-file.txt")

	// copy the file
	createdFile, byteCount, err := fileutil.CopyFile(sourceFilePath, targetFilePath, true)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("Copied %v bytes to %v.\n", byteCount, filepath.Base(createdFile))
	}

	// Output: Copied 13 bytes to target-file.txt.
}

func ExampleCopyFile_toDirectory() {
	// create a file that we will copy (usually this already exists)
	sourceFilePath := filepath.Join(os.TempDir(), "source-file.txt")

	err := os.WriteFile(sourceFilePath, []byte("Hello, World!"), 0600)
	if err != nil {
		fmt.Println(err)
	}

	// define the target path to copy the file to
	targetDirectory, _ := os.MkdirTemp("", "target-directory-*")

	// copy the file
	createdFile, byteCount, err := fileutil.CopyFile(sourceFilePath, targetDirectory, true)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("Copied %v bytes to %v.\n", byteCount, filepath.Base(createdFile))
	}

	// Output: Copied 13 bytes to source-file.txt.
}
