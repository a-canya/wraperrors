package wraperrors

import "errors"

type wrapError struct {
	wrapper error
	wrapped error
}

// New creates a new wrap error
func New(wrapper, wrapped error) error {
	return &wrapError{
		wrapper: wrapper,
		wrapped: wrapped,
	}
}

// Error returns a string describing the wrap error and allows warp errors to implement the error interface
func (err *wrapError) Error() string {
	return err.wrapper.Error() + ", " + err.wrapped.Error()
}

// Is compares the equality of a warp error to a target error. This method allows wrap errors to implement ths errors Is
// interface, and allows wrap errors to be compared using `errors.Is(myWrapError, anyOtherError)`
//
// A wrap error is equal to a target error if:
//
// (a) the wrap error is exactly equal (== operator) to the target error,
//
// (b) the wrapper error is equal to the target error using the errors.Is function, or
//
// (c) the wrapped error is equal to the target error using the errors.Is function
//
// (notice that b and c differ on the word wrapper/wrapped)
func (err *wrapError) Is(target error) bool {
	return err == target || errors.Is(err.wrapper, target) || errors.Is(err.wrapped, target)
}

// Unwrap returns the wrapped error
func (err *wrapError) Unwrap() error {
	return err.wrapped
}
