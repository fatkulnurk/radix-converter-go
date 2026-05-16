package radixerrors_test

import (
	"errors"
	"testing"

	"github.com/fatkulnurk/radix-converter-go/radixerrors"
)

func TestError_IsErrorType(t *testing.T) {
	var err error = radixerrors.ErrNegativeInput
	var target *radixerrors.Error
	if !errors.As(err, &target) {
		t.Fatal("ErrNegativeInput is not a *radixerrors.Error")
	}
}

func TestError_MessageContainsInfo(t *testing.T) {
	err := radixerrors.NewInvalidCharError('!')
	msg := err.Error()
	if msg == "" {
		t.Error("Error message is empty")
	}
	// Should contain the invalid char.
	if !containsRune(msg, '!') {
		t.Errorf("Error message %q does not contain the invalid char", msg)
	}
}

func TestError_UnknownConverter(t *testing.T) {
	err := radixerrors.NewUnknownConverterError("foo")
	msg := err.Error()
	if !contains(msg, "foo") {
		t.Errorf("Error message %q does not contain converter name \"foo\"", msg)
	}
}

func TestError_Wrapping(t *testing.T) {
	err := radixerrors.ErrEmptyEncoded
	var target *radixerrors.Error
	if !errors.As(err, &target) {
		t.Fatal("ErrEmptyEncoded cannot be unwrapped as *radixerrors.Error")
	}
}

func TestError_ErrorsIs(t *testing.T) {
	// Errors of type *radixerrors.Error should match each other via Is.
	err1 := radixerrors.ErrNegativeInput
	err2 := radixerrors.ErrNegativeInput
	if !errors.Is(err1, err2) {
		t.Error("ErrNegativeInput should be equal to itself via errors.Is")
	}
}

func contains(s string, substr string) bool {
	return len(s) > 0 && len(substr) > 0 && indexOf(s, substr) >= 0
}

func indexOf(s, substr string) int {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return i
		}
	}
	return -1
}

func containsRune(s string, r rune) bool {
	for _, c := range s {
		if c == r {
			return true
		}
	}
	return false
}
