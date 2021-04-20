package wraperrors

/*
wraperrors provides a simple interface to work with wrap errors. A wrap error has two errors:
- A wrapper error
- A wrapped error

The wrapper error is suposed to be generated by a higher level part of the code and provide more general information,
whereas the wrapped error is supposed to have more specific details about the actual error.

The package also provides a SimpleError type which is useful to create wrap errors that will have the simple error as
the wrapper.

You can create a new wrap error:

```go
const myErr = wraperrors.SimpleError("something went wrong")
err1 := someFunc()
err2 := wraperrors.New(myErr, err1)
```

A wrap error has a wrapper error and a wrapped error. Its main difference with the standard errors.Error
 `fmt.Errorf("something went wrong, %w", err1)` because it allows easy comparisons
using the Is error comparer:

```go
const (
	TemporalError = wraperrors.SimpleError("temporal error, try later")
	FatalError    = wraperrors.SimpleError("fatal error")
)

func f() error {
	err := getSomethingFromTheInternet()
	if err == ConnectionTimeOut {
		return TemporalError.Wrap(err)
	} else if err != nil {
		return FatalError.Wrap(err)
	}
	return nil
}

err := f()
if errors.Is(err, TemporalError) {
	// retry after timeout
} else {
	fmt.Printf("%s", err.Error())
	// logs fatal error and the error message returned by getSomethingFromTheInternet
}
```

Instead of wraperrors.SimpleError, other errors can also be used as the wrapper error.


*/
