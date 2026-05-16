package radixerrors

import "fmt"

// Error represents a radix converter error.
type Error struct {
	Op  string // operation being performed (e.g. "encode", "decode")
	Err error  // underlying error
}

func (e *Error) Error() string {
	return fmt.Sprintf("radix converter: %s: %v", e.Op, e.Err)
}

func (e *Error) Unwrap() error {
	return e.Err
}

// Is implements error matching for Error.
func (e *Error) Is(target error) bool {
	_, ok := target.(*Error)
	return ok
}

var (
	// ErrNegativeInput is returned when attempting to encode a negative number.
	ErrNegativeInput = &Error{Op: "encode", Err: fmt.Errorf("input number must be non-negative")}

	// ErrEmptyEncoded is returned when attempting to decode an empty string.
	ErrEmptyEncoded = &Error{Op: "decode", Err: fmt.Errorf("encoded value cannot be empty")}

	// ErrInvalidChar is returned when an invalid character is found during decoding.
	ErrInvalidChar = &Error{Op: "decode", Err: fmt.Errorf("invalid character found")}

	// ErrUnknownConverter is returned when a converter type is not recognized.
	ErrUnknownConverter = &Error{Op: "create", Err: fmt.Errorf("unknown converter type")}

	// ErrConverterAlreadyRegistered is returned when registering a duplicate custom converter.
	ErrConverterAlreadyRegistered = &Error{Op: "register", Err: fmt.Errorf("converter is already registered")}

	// ErrConverterNotRegistered is returned when getting a converter that doesn't exist.
	ErrConverterNotRegistered = &Error{Op: "get", Err: fmt.Errorf("converter is not registered")}
)

// NewInvalidCharError creates an Error with the invalid character included in the message.
func NewInvalidCharError(char rune) error {
	return &Error{Op: "decode", Err: fmt.Errorf("invalid character: %q", char)}
}

// NewUnknownConverterError creates an Error with the unknown type name.
func NewUnknownConverterError(name string) error {
	return &Error{Op: "create", Err: fmt.Errorf("unknown converter type: %s", name)}
}
