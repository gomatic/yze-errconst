package a

import (
	"errors"
	"fmt"
)

// errSentinel is created with errors.New and must be flagged.
var errSentinel = errors.New("boom") // want `use a sentinel error constant \(errs\.Const\) instead of errors\.New`

// wrap uses fmt.Errorf with a trailing %w and is allowed.
func wrap(cause error) error {
	return fmt.Errorf("context: %w", cause)
}

// wrapLeading wraps an injected cause with a leading %w (the canonical
// sentinel-leading shape). The v1 %w proxy allows this: it deliberately does not
// trace the wrapped argument, so a non-sentinel error value is accepted exactly
// as the convention permits ("adding context to a sentinel or an injected cause").
func wrapLeading(cause error) error {
	return fmt.Errorf("%w: extra context", cause)
}

// noWrap uses fmt.Errorf with no %w and must be flagged.
func noWrap() error {
	return fmt.Errorf(
		"plain failure",
	) // want `use a sentinel error constant, or wrap a cause with %w, instead of fmt\.Errorf`
}

// dynamicFormat uses a non-literal format; it cannot be statically judged and is
// not flagged.
func dynamicFormat(format string, cause error) error {
	return fmt.Errorf(format, cause)
}

// plainFormat is a constant format string that does not wrap a cause.
const plainFormat = "plain failure %d"

// constFormat passes a const-string format to fmt.Errorf. The v1 %w proxy only
// inspects a literal format argument, so a const-held format escapes the check
// and is not flagged — a deliberate, documented escape hatch.
func constFormat() error {
	return fmt.Errorf(plainFormat, 1)
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
