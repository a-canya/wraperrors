package wraperrors

// SimpleError returns a simple error useful to define constant simple errors that will wrap other errors.
//
// ```go
// const MyError = wraperrors.SimpleError("my error message")
// ```
type SimpleError string

// Error returns the simple error's string and allows simple errors to implement the error interface
func (err SimpleError) Error() string {
	return string(err)
}

// Is compares the equality of the simple error and a target
func (err SimpleError) Is(target error) bool {
	return err == target
}

// Wrap creates a wrap error that has this simple error as wrapper error and wrapped as the wrapped error
func (err SimpleError) Wrap(wrapped error) error {
	return New(err, wrapped)
}
