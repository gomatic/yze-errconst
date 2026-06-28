package a

import (
	"errors"
	"fmt"
)

// errSentinel is created with errors.New and must be flagged.
var errSentinel = errors.New("boom") // want `errors.New`

// wrap uses fmt.Errorf with %w and is allowed.
func wrap(cause error) error {
	return fmt.Errorf("context: %w", cause)
}

// noWrap uses fmt.Errorf with no %w and must be flagged.
func noWrap() error {
	return fmt.Errorf("plain failure") // want `fmt.Errorf`
}

// dynamicFormat uses a non-literal format; it cannot be statically judged and is
// not flagged.
func dynamicFormat(format string, cause error) error {
	return fmt.Errorf(format, cause)
}

// dynamicCall invokes a func value, exercising the non-static-callee path.
func dynamicCall() {
	f := func() {}
	f()
}

// benign calls a non-constructor stdlib function and must not be flagged.
func benign() string {
	return fmt.Sprintf("hi %d", 1)
}
