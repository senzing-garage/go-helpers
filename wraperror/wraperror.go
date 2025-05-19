package wraperror

import (
	"encoding/json"
	"fmt"
	"regexp"
	"runtime"
)

const (
	NoMessage  = ""
	callerSkip = 2
)

// ----------------------------------------------------------------------------
// Public functions
// ----------------------------------------------------------------------------

/*
The Error function returns a wrapped error, if err != nil.

Input
  - err: The unwrapped/raw error.

Output
  - Either nil, or a wrapped error.
*/
func Error(err error) error {
	var (
		errFormat    string
		functionName string
	)

	if err != nil {
		functionName = findFunctionName(callerSkip)

		switch {
		case len(functionName) > 0:
			errFormat = fmt.Sprintf(
				`{"function": %s, "error": %s}`,
				formatFunction(),
				formatError(err),
			)

			return fmt.Errorf(errFormat, functionName, err) //nolint
		default:
			errFormat = fmt.Sprintf(
				`{"error": %s}`,
				formatError(err),
			)

			return fmt.Errorf(errFormat, err) //nolint
		}
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
		text         string
	)

	if err != nil {
		text = fmt.Sprintf(format, messages...)
		functionName = findFunctionName(callerSkip)

		switch {
		case len(functionName) > 0 && len(text) > 0:
			errFormat = fmt.Sprintf(
				`{"function": %s, "text": %s, "error": %s}`,
				formatFunction(),
				formatText(text),
				formatError(err),
			)

			return fmt.Errorf(errFormat, functionName, text, err) //nolint
		case len(functionName) > 0:
			errFormat = fmt.Sprintf(
				`{"function": %s, "error": %s}`,
				formatFunction(),
				formatError(err),
			)

			return fmt.Errorf(errFormat, functionName, err) //nolint
		case len(text) > 0:
			errFormat = fmt.Sprintf(
				`{"text": %s, "error": %s}`,
				formatText(text),
				formatError(err),
			)

			return fmt.Errorf(errFormat, text, err) //nolint
		default:
			errFormat = fmt.Sprintf(
				`{"error": %s}`,
				formatError(err),
			)

			return fmt.Errorf(errFormat, err) //nolint
		}
	}

	return nil
}

// ----------------------------------------------------------------------------
// Private functions
// ----------------------------------------------------------------------------

func findFunctionName(callerSkip int) string {
	var result string

	pc, _, _, ok := runtime.Caller(callerSkip)
	if ok {
		callingFunction := runtime.FuncForPC(pc)
		runtimeFunc := regexp.MustCompile(`([^/]+$)`)
		result = runtimeFunc.FindString(callingFunction.Name())
	}

	return result
}

func formatError(err error) string {
	result := `"%w"`
	if isJSON(err.Error()) {
		result = `%w`
	}

	return result
}

func formatFunction() string {
	return `"%s"`
}

func formatText(text string) string {
	result := `"%s"`
	if isJSON(text) {
		result = `%s`
	}

	return result
}

func isJSON(unknownText string) bool {
	return json.Valid([]byte(unknownText))
}
