package wraperror

import (
	"encoding/json"
	"fmt"
	"regexp"
	"runtime"
)

/*
The Errorf function returns a wrapped error, if err != nil.

Input
  - err: The unwrapped/raw error.
  - format: The format string (think fmt.Sprintf())
  - messages: values to be put into the format string.

Output
  - Either nil, or a wrapped error.
*/
func Error(err error) error {
	var (
		errFormat    string
		functionName string
	)

	if err != nil {

		// Result if function name is available.

		pc, _, _, ok := runtime.Caller(1)
		if ok {
			callingFunction := runtime.FuncForPC(pc)
			runtimeFunc := regexp.MustCompile(`([^/]+$)`)
			functionName = runtimeFunc.FindString(callingFunction.Name())
			errFormat = `{"function": "%s", "error": "%w"}`
			if isJSON(err.Error()) {
				errFormat = `{"function": "%s", "error": %w}`
			}

			return fmt.Errorf(errFormat, functionName, err)
		}

		// Default JSON.

		errFormat = `{"error": "%w"}`
		if isJSON(err.Error()) {
			errFormat = `{"error": %w}`
		}

		return fmt.Errorf(errFormat, err)
	}

	return nil
}

/*
The Errorf function returns a wrapped error, if err != nil.

Input
  - err: The unwrapped/raw error.
  - format: The format string (think fmt.Sprintf())
  - messages: values to be put into the format string.

Output
  - Either nil, or a wrapped error.
*/
func Errorf(err error, format string, messages ...any) error {
	var (
		errFormat    string
		functionName string
	)

	if err != nil {
		text := fmt.Sprintf(format, messages...)

		textFormat := `"%s"`
		if isJSON(text) {
			textFormat = `%s`
		}

		errorFormat := `"%w"`
		if isJSON(err.Error()) {
			errorFormat = `%w`
		}

		functionFormat := `%s`

		// Result if function name is available.

		pc, _, _, ok := runtime.Caller(1)
		if ok {
			callingFunction := runtime.FuncForPC(pc)
			runtimeFunc := regexp.MustCompile(`([^/]+$)`)
			functionName = runtimeFunc.FindString(callingFunction.Name())
			errFormat = fmt.Sprintf(`{"function": %s, "text": %s, "error": %s}`, functionFormat, textFormat, errorFormat)
			return fmt.Errorf(errFormat, functionName, text, err)
		}

		// Default JSON.

		errFormat = fmt.Sprintf(`{"text": %s, "error": %s}`, textFormat, errorFormat)


FIXME:  Start here.

		return fmt.Errorf(errFormat, text, err)
	}

	return nil
}

func Errorf1(err error, format string, messages ...any) error {
	if err != nil {

		pc, _, _, ok := runtime.Caller(1)
		if ok {
			callingFunction := runtime.FuncForPC(pc)
			runtimeFunc := regexp.MustCompile(`([^/]+$)`)
			functionName := runtimeFunc.FindString(callingFunction.Name())
			format = functionName + "(): " + format
		}

		errFormat := format + " >>> %w"
		messages = append(messages, err)

		return fmt.Errorf(errFormat, messages...) //nolint:err113
	}

	return nil
}

func Errorf2(err error, format string, messages ...any) error {
	var functionName string

	if err != nil {

		text := fmt.Sprintf(format, messages...)

		pc, file, line, ok := runtime.Caller(1)
		if ok {
			callingFunction := runtime.FuncForPC(pc)
			runtimeFunc := regexp.MustCompile(`([^/]+$)`)
			functionName = runtimeFunc.FindString(callingFunction.Name())
			format := `{"function": "%s", "text": "%s", "file": "%s", "line": %d, "error": "%w"}`
			if isJSON(err.Error()) {
				format = `{"function": "%s", "text": "%s", "file": "%s", "line": %d, "error": %w}`
			}

			return fmt.Errorf(format, functionName, text, file, line, err)

		}

	}

	return nil
}

func isJSON(unknownText string) bool {
	return json.Valid([]byte(unknownText))
}
