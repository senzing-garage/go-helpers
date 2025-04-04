package main

import (
	"os"
	"testing"
)

func TestMain(test *testing.T) {
	test.Parallel()

	_ = test

	main()
}

func TestMain_withArgs(test *testing.T) {
	test.Parallel()

	_ = test
	os.Args = []string{"command-name", "-j", "{}"}

	main()
}
