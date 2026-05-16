package radixconverter_test

import (
	"errors"
	"testing"

	"github.com/fatkulnurk/radix-converter-go"
)

func TestError_IsErrorType(t *testing.T) {
	var err error = radixconverter.ErrNegativeInput
	var target *radixconverter.Error
	if !errors.As(err, &target) {
		t.Fatal("ErrNegativeInput is not a *radixconverter.Error")
	}
}

func TestError_MessageContainsInfo(t *testing.T) {
	err := radixconverter.NewInvalidCharError('!')
	msg := err.Error()
	if msg == "" {
		t.Error("Error message is empty")
	}
	if !containsRune(msg, '!') {
		t.Errorf("Error message %q does not contain the invalid char", msg)
	}
}

func TestError_UnknownConverter(t *testing.T) {
	err := radixconverter.NewUnknownConverterError("foo")
	msg := err.Error()
	if !contains(msg, "foo") {
		t.Errorf("Error message %q does not contain converter name \"foo\"", msg)
	}
}

func TestError_Wrapping(t *testing.T) {
	err := radixconverter.ErrEmptyEncoded
	var target *radixconverter.Error
	if !errors.As(err, &target) {
		t.Fatal("ErrEmptyEncoded cannot be unwrapped as *radixconverter.Error")
	}
}

func TestError_ErrorsIs(t *testing.T) {
	err1 := radixconverter.ErrNegativeInput
	err2 := radixconverter.ErrNegativeInput
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
