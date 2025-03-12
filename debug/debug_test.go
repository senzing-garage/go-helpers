package debug

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// ----------------------------------------------------------------------------
// Test interface methods
// ----------------------------------------------------------------------------

func TestDebug_Print_0(test *testing.T) {
	debugger, err := New(0)
	testError(test, err)
	debugger.Print(1, "Should not print level 1\n")
	debugger.Print(4, "Should not print level 4\n")
	debugger.Print(5, "Should not print level 5\n")
	debugger.Print(9, "Should not print level 9\n")
}

func TestDebug_Print_5(test *testing.T) {
	debugger, err := New(5)
	testError(test, err)
	debugger.Print(1, "Should not print level 1\n")
	debugger.Print(4, "Should not print level 4\n")
	debugger.Print(5, "Should print level 5\n")
	debugger.Print(9, "Should print level 9\n")
}

func TestDebug_Print_10(test *testing.T) {
	debugger, err := New(10)
	testError(test, err)
	debugger.Print(1, "Should not print level 1\n")
	debugger.Print(4, "Should not print level 4\n")
	debugger.Print(5, "Should not print level 5\n")
	debugger.Print(9, "Should not print level 9\n")
}

func TestDebug_Printf_0(test *testing.T) {
	debugger, err := New(0)
	testError(test, err)
	debugger.Printf(1, "Should not print level %d\n", 1)
	debugger.Printf(4, "Should not print level %d\n", 4)
	debugger.Printf(5, "Should not print level %d\n", 5)
	debugger.Printf(9, "Should not print level %d\n", 9)
}

func TestDebug_Printf_5(test *testing.T) {
	debugger, err := New(5)
	testError(test, err)
	debugger.Printf(1, "Should not print level %d\n", 1)
	debugger.Printf(4, "Should not print level %d\n", 4)
	debugger.Printf(5, "Should print level %d\n", 5)
	debugger.Printf(9, "Should print level %d\n", 9)
}

func TestDebug_Printf_10(test *testing.T) {
	debugger, err := New(10)
	testError(test, err)
	debugger.Printf(1, "Should not print level %d\n", 1)
	debugger.Printf(4, "Should not print level %d\n", 4)
	debugger.Printf(5, "Should not print level %d\n", 5)
	debugger.Printf(9, "Should not print level %d\n", 9)
}

func TestDebug_Println_0(test *testing.T) {
	debugger, err := New(0)
	testError(test, err)
	debugger.Println(1, "Should not print level 1")
	debugger.Println(4, "Should not print level 4")
	debugger.Println(5, "Should not print level 5")
	debugger.Println(9, "Should not print level 9")
}

func TestDebug_Println_5(test *testing.T) {
	debugger, err := New(5)
	testError(test, err)
	debugger.Println(1, "Should not print level 1")
	debugger.Println(4, "Should not print level 4")
	debugger.Println(5, "Should print level 5")
	debugger.Println(9, "Should print level 9")
}

func TestDebug_Println_10(test *testing.T) {
	debugger, err := New(10)
	testError(test, err)
	debugger.Println(1, "Should not print level 1")
	debugger.Println(4, "Should not print level 4")
	debugger.Println(5, "Should not print level 5")
	debugger.Println(9, "Should not print level 9")
}

func testError(test *testing.T, err error) {
	if err != nil {
		assert.FailNow(test, err.Error())
	}
}
