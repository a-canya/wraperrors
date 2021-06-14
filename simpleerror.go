package wraperrors

// SimpleError is a simple string error. It is specially useful to define constant errors and to wrap another other to
// create a "wrap error".
//
// ```go
// const MyError = wraperrors.SimpleError("something went wrong")
// var wrapErr = wraperrors.SimpleError("program will shut down").Wrap(MyError)
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
