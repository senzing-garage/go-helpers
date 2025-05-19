package fileutil_test

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"testing"

	"github.com/senzing-garage/go-helpers/fileutil"
	"github.com/senzing-garage/go-helpers/wraperror"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var errForPackage = errors.New("fileutil_test")

// ----------------------------------------------------------------------------
// Test CopyFile() function
// ----------------------------------------------------------------------------

func TestCopyFile_Basic1(test *testing.T) {
	test.Parallel()

	destinationDir := destinationDirectoryPath()
	destinationFile := filepath.Join(destinationDir, "basic_file_1.txt")
	sourceFile, fileSize := sourceFilePath1()
	createdFile, byteCount, err := fileutil.CopyFile(sourceFile, destinationFile, true)
	require.NoError(test, err)
	assert.Equal(test, fileSize, byteCount, "Byte Count for CopyFile() not as expected for basic file 1")
	assert.Equal(test, destinationFile, createdFile, "Created file path is not as expected for basic file 1")

	stat, err := os.Stat(destinationFile)
	require.NoError(test, err)
	assert.Equal(test, fileSize, stat.Size(), "File size of basic file 1 not as expected post CopyFile()")

	content, err := os.ReadFile(sourceFile)
	require.NoError(test, err)

	expectedContent := string(content)

	content, err = os.ReadFile(destinationFile)
	require.NoError(test, err)

	actualContent := string(content)
	assert.Equal(test, expectedContent, actualContent, "File contents of basic file 1 not as expected post CopyFile()")
}

func TestCopyFile_Basic2(test *testing.T) {
	test.Parallel()

	destinationDir := destinationDirectoryPath()
	destinationFile := filepath.Join(destinationDir, "basic_file_2.txt")
	sourceFile, fileSize := sourceFilePath2()
	createdFile, byteCount, err := fileutil.CopyFile(sourceFile, destinationFile, false)
	require.NoError(test, err)
	assert.Equal(test, fileSize, byteCount, "Byte Count for CopyFile() not as expected for basic file 2")
	assert.Equal(test, destinationFile, createdFile, "Created file path is not as expected for basic file 2")

	stat, err := os.Stat(destinationFile)
	require.NoError(test, err)
	assert.Equal(test, fileSize, stat.Size(), "File size of basic file 2 not as expected post CopyFile()")

	content, err := os.ReadFile(sourceFile)
	require.NoError(test, err)

	expectedContent := string(content)

	content, err = os.ReadFile(destinationFile)
	require.NoError(test, err)

	actualContent := string(content)
	assert.Equal(test, expectedContent, actualContent, "File contents of basic file 2 not as expected post CopyFile()")
}

func TestCopyFile_ToDirectory(test *testing.T) { //nolint
	// test.Parallel()
	destinationDir := destinationDirectoryPath()
	sourceFile, fileSize := sourceFilePath1()

	// determine what the file name should be
	destinationFile := filepath.Join(destinationDir, filepath.Base(sourceFile))

	createdFile, byteCount, err := fileutil.CopyFile(sourceFile, destinationDir, true)
	require.NoError(test, err)
	assert.Equal(test, fileSize, byteCount, "Byte Count for CopyFile() not as expected when copying to directory")
	assert.Equal(test, destinationFile, createdFile, "Created file path is not as expected for CopyFile() to directory")

	stat, err := os.Stat(destinationFile)
	require.NoError(test, err)
	assert.Equal(test, fileSize, stat.Size(), "File size not as expected post CopyFile() when copying to directory")

	content, err := os.ReadFile(sourceFile)
	require.NoError(test, err)

	expectedContent := string(content)

	content, err = os.ReadFile(destinationFile)
	require.NoError(test, err)

	actualContent := string(content)
	assert.Equal(test, expectedContent, actualContent, "File contents of not as expected post CopyFile() to directory")
}

func TestCopyFile_WithOverwrite(test *testing.T) {
	test.Parallel()

	destinationDir := destinationDirectoryPath()
	destinationFile := filepath.Join(destinationDir, "with_overwrite.txt")

	expectedContent := "Already Exists"
	_, err := createTextFile(destinationFile, expectedContent)
	require.NoError(test, err)

	sourceFile, fileSize := sourceFilePath2()

	content, err := os.ReadFile(sourceFile)
	require.NoError(test, err)

	expectedContent = string(content)

	createdFile, byteCount, err := fileutil.CopyFile(sourceFile, destinationFile, true)
	require.NoError(test, err)
	assert.Equal(test, fileSize, byteCount, "Byte Count for CopyFile() not as expected for overwritten file")
	assert.Equal(test, destinationFile, createdFile, "Overwritten file path is not as expected for CopyFile()")

	stat, err := os.Stat(destinationFile)
	require.NoError(test, err)
	assert.Equal(test, fileSize, stat.Size(), "File size not as expected post CopyFile() for overwritten file")

	content, err = os.ReadFile(destinationFile)
	require.NoError(test, err)

	actualContent := string(content)
	assert.Equal(
		test,
		expectedContent,
		actualContent,
		"File contents of overwritten file not as expected post CopyFile()",
	)
}

func TestCopyFile_NoOverwrite(test *testing.T) {
	test.Parallel()

	destinationDir := destinationDirectoryPath()
	destinationFile := filepath.Join(destinationDir, "no_overwrite.txt")

	expectedContent := "Already Exists again"
	byteCount, err := createTextFile(destinationFile, expectedContent)
	require.NoError(test, err)

	sourceFile, _ := sourceFilePath1()
	_, _, err = fileutil.CopyFile(sourceFile, destinationFile, false)
	require.Error(test, err, "Expected an error when attempting to overwrite file with CopyFile()")

	stat, err := os.Stat(destinationFile)
	require.NoError(test, err)
	assert.Equal(test, byteCount, stat.Size(), "File size not as expected post CopyFile() with no overwrite")

	content, err := os.ReadFile(destinationFile)
	require.NoError(test, err)

	actualContent := string(content)
	assert.Equal(test, expectedContent, actualContent,
		"File contents not as expected post CopyFile() with no overwrite")
}

func TestCopyFile_ToDirectoryWithOverwrite(test *testing.T) {
	test.Parallel()

	destinationDir := destinationDirectoryPath()
	sourceFile, fileSize := sourceFilePath1()

	// determine what the file name should be
	destinationFile := filepath.Join(destinationDir, filepath.Base(sourceFile))

	// remove the destination file if it already exists
	_, err := os.Stat(destinationFile)
	if err == nil {
		// remove the file
		err := os.Remove(destinationFile)
		require.NoError(test, err)
	} else if !errors.Is(err, fs.ErrNotExist) {
		// file exists, but we got a different error
		require.NoError(test, err)
	}

	_, err = createTextFile(destinationFile, "Already Exists")
	require.NoError(test, err)

	content, err := os.ReadFile(sourceFile)
	require.NoError(test, err)

	expectedContent := string(content)

	createdFile, byteCount, err := fileutil.CopyFile(sourceFile, destinationDir, true)
	require.NoError(test, err)
	assert.Equal(
		test,
		fileSize,
		byteCount,
		"Byte Count for CopyFile() to directory not as expected for overwritten file",
	)
	assert.Equal(
		test,
		destinationFile,
		createdFile,
		"Overwritten file path is not as expected for CopyFile() to directory",
	)

	stat, err := os.Stat(destinationFile)
	require.NoError(test, err)
	assert.Equal(
		test,
		fileSize,
		stat.Size(),
		"File size of overwritten file not as expected post CopyFile() to directory",
	)

	content, err = os.ReadFile(destinationFile)
	require.NoError(test, err)

	actualContent := string(content)
	assert.Equal(
		test,
		expectedContent,
		actualContent,
		"File contents of overwritten file not as expected post CopyFile() to directory",
	)
}

func TestCopyFile_ToDirectoryNoOverwrite(test *testing.T) {
	test.Parallel()

	destinationDir := destinationDirectoryPath()
	sourceFile, _ := sourceFilePath2()

	// determine what the file name should be
	destinationFile := filepath.Join(destinationDir, filepath.Base(sourceFile))

	// remove the destination file if it already exists
	_, err := os.Stat(destinationFile)
	if err == nil {
		// remove the file
		err := os.Remove(destinationFile)
		require.NoError(test, err)
	} else if !errors.Is(err, fs.ErrNotExist) {
		// file exists, but we got a different error
		require.NoError(test, err)
	}

	expectedContent := "Already Exists"
	byteCount, err := createTextFile(destinationFile, expectedContent)
	require.NoError(test, err)

	_, _, err = fileutil.CopyFile(sourceFile, destinationDir, false)
	require.Error(test, err, "Expected an error when attempting to overwrite file with CopyFile() to directory")

	stat, err := os.Stat(destinationFile)
	require.NoError(test, err)
	assert.Equal(
		test,
		byteCount,
		stat.Size(),
		"File size not as expected post CopyFile() to directory with no overwrite",
	)

	content, err := os.ReadFile(destinationFile)
	require.NoError(test, err)

	actualContent := string(content)
	assert.Equal(test, expectedContent, actualContent,
		"File contents not as expected post CopyFile() to directory with no overwrite")
}

func TestCopyFile_FromDirectory(test *testing.T) {
	test.Parallel()

	sourceDir := sourceDirectoryPath()
	destinationDir := destinationDirectoryPath()
	destinationFile := filepath.Join(destinationDir, "directory_copy")
	_, _, err := fileutil.CopyFile(sourceDir, destinationFile, true)
	require.Error(test, err, "Did not get expected error when trying to copy a directory")
}

func TestCopyFile_SourceNotFound(test *testing.T) {
	test.Parallel()

	sourceDir := sourceDirectoryPath()
	sourceFile := filepath.Join(sourceDir, "does_not_exist.txt")
	destinationDir := destinationDirectoryPath()
	destinationFile := filepath.Join(destinationDir, "will_not_exist.txt")
	_, _, err := fileutil.CopyFile(sourceFile, destinationFile, true)
	require.Error(test, err, "Did not get expected error when trying to copy a non-existent file")
}

func TestCopyFile_DestinationNotFound(test *testing.T) {
	test.Parallel()

	sourceFile, _ := sourceFilePath1()
	destinationDir := destinationDirectoryPath()
	badSubDirectory := filepath.Join(destinationDir, "does_not_exist")
	destinationFile := filepath.Join(badSubDirectory, "will_not_exist.txt")
	_, _, err := fileutil.CopyFile(sourceFile, destinationFile, true)
	require.Error(test, err, "Did not get expected error when trying to copy a bad destination path")
}

// ----------------------------------------------------------------------------
// Private functions
// ----------------------------------------------------------------------------

func baseDirectoryPath() string {
	return filepath.FromSlash("../target/test/fileutil")
}

func sourceDirectoryPath() string {
	return filepath.Join(baseDirectoryPath(), "source")
}

func destinationDirectoryPath() string {
	return filepath.Join(baseDirectoryPath(), "destination")
}

func sourceFilePath1() (string, int64) {
	return filepath.Join(sourceDirectoryPath(), "Five_Byte_File.txt"), 5
}

func sourceFilePath2() (string, int64) {
	return filepath.Join(sourceDirectoryPath(), "Ten_Byte_File.txt"), 10
}

func createTextFile(path string, text string) (int64, error) {
	source, err := os.Create(filepath.Clean(path))
	if err != nil {
		return 0, wraperror.Errorf(err, "failed to create file (%v)", path)
	}

	defer source.Close()

	byteCount, err := source.WriteString(text)
	if err != nil {
		return 0, wraperror.Errorf(err, "failed to write content (%v) to file (%v)", text, path)
	}

	return int64(byteCount), err
}

func createTextFileN(path string, byteCount int64) (int64, error) {
	var writeCount int64

	source, err := os.Create(filepath.Clean(path))
	if err != nil {
		return writeCount, wraperror.Errorf(err, "failed to create file (%v)", path)
	}

	defer source.Close()

	for index := range byteCount {
		count, err := source.WriteString("A")
		if err != nil {
			return writeCount, wraperror.Errorf(err, "failed to write letter (%v) to file (%v)",
				index, path)
		}

		writeCount += int64(count)
	}

	if writeCount != byteCount {
		return writeCount, wraperror.Errorf(errForPackage, "wrote wrong number of bytes (%v) to file (%v) error",
			writeCount, path)
	}

	return writeCount, wraperror.Errorf(err, "fileutil.createTextFileN error")
}

func outputln(message ...any) {
	fmt.Println(message...) //nolint
}

func useBase64(thing int64) {
	_ = thing
}

// ----------------------------------------------------------------------------
// Test harness
// ----------------------------------------------------------------------------

func TestMain(m *testing.M) {
	err := setup()
	if err != nil {
		outputln(err)
		os.Exit(1)
	}

	code := m.Run()

	err = teardown()
	if err != nil {
		outputln(err)
	}

	os.Exit(code)
}

func setup() error {
	baseDir := baseDirectoryPath()

	// remove any previously existing test directory
	err := os.RemoveAll(baseDir)
	if err != nil {
		return wraperror.Errorf(err, "failed to delete old test targets in %v", baseDir)
	}

	// define the source and destination directories
	sourceDir := sourceDirectoryPath()
	destinationDir := destinationDirectoryPath()

	// make the source directory and any required parents
	err = os.MkdirAll(sourceDir, 0o770)
	if err != nil {
		return wraperror.Errorf(err, "failed to create source directory (%v)", sourceDir)
	}

	// make the destination directory and any required parents
	err = os.MkdirAll(destinationDir, 0o770)
	if err != nil {
		return wraperror.Errorf(err, "failed to create destination directory (%v)", destinationDir)
	}

	// define paths to
	sourcePath1, fileSize1 := sourceFilePath1()
	sourcePath2, fileSize2 := sourceFilePath2()

	byteCount, err := createTextFileN(sourcePath1, fileSize1)
	if err != nil {
		return err
	}

	useBase64(byteCount)

	byteCount, err = createTextFileN(sourcePath2, fileSize2)
	if err != nil {
		return err
	}

	useBase64(byteCount)

	return err
}

func teardown() error {
	baseDir := baseDirectoryPath()

	// remove any previously existing test directory
	err := os.RemoveAll(baseDir)
	if err != nil {
		return wraperror.Errorf(err, "failed to delete old test targets in %v", baseDir)
	}

	return wraperror.Errorf(err, "fileutil.teardown error")
}
