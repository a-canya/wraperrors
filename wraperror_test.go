package wraperrors_test

import (
	"errors"
	"testing"

	"github.com/a-canya/wraperrors"
)

const (
	wrappedErr wraperrors.SimpleError = "I failed at doing something"
	wrapperErr wraperrors.SimpleError = "f() failed, program will finish"
)

func f() error {
	return wraperrors.SimpleError(wrappedErr)
}

func g() error {
	err := f()
	if err != nil {
		return wraperrors.SimpleError(wrapperErr).Wrap(err)
	}

	return nil
}

func TestErrIsWrappedErr(t *testing.T) {
	err := g()

	if !errors.Is(err, wrappedErr) {
		t.Fail()
	}
}

func TestWrappedErrIsNotErr(t *testing.T) {
	err := g()

	if errors.Is(wrappedErr, err) {
		t.Fail()
	}
}

func TestErrIsWrapperErr(t *testing.T) {
	err := g()

	if !errors.Is(err, wrapperErr) {
		t.Fail()
	}
}

func TestWrapperErrIsNotErr(t *testing.T) {
	err := g()

	if errors.Is(wrapperErr, err) {
		t.Fail()
	}
}
