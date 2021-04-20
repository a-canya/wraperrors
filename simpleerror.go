package wraperrors

type simpleError string

// SimpleError returns a simple error useful to define constant simple errors that will wrap other errors.
//
// ```go
// const MyError = wraperrors.simpleError("my error message")
// ```
func SimpleError(str string) error {
	return simpleError(str)
}

func (err simpleError) Error() string {
	return string(err)
}

// Is compares equality of err and target
func (err simpleError) Is(target error) bool {
	return err == target
}

// Wrap wraps the wrappedErr into err, creating an error that has err as its wrapper error and wrappedErr as its wrapped
// error
func (err simpleError) Wrap(wrapped error) error {
	return New(err, wrapped)
}
