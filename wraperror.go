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

func (err *wrapError) Error() string {
	return err.wrapper.Error() + ", " + err.wrapped.Error()
}

// Is compares the equality of err and target, or the equality of the wrapper error in err with target, or the equality
// of the wrapped error in err and target
func (err *wrapError) Is(target error) bool {
	return err == target || errors.Is(err.wrapper, target) || errors.Is(err.wrapped, target)
}

// Unwrap unwraps the wrapped error
func (err *wrapError) Unwrap() error {
	return err.wrapped
}
