package radixconverter

import "fmt"

// Error represents a radix converter error.
type Error struct {
	Op  string
	Err error
}

func (e *Error) Error() string {
	return fmt.Sprintf("radix converter: %s: %v", e.Op, e.Err)
}

func (e *Error) Unwrap() error {
	return e.Err
}

func (e *Error) Is(target error) bool {
	_, ok := target.(*Error)
	return ok
}

var (
	ErrNegativeInput             = &Error{Op: "encode", Err: fmt.Errorf("input number must be non-negative")}
	ErrEmptyEncoded              = &Error{Op: "decode", Err: fmt.Errorf("encoded value cannot be empty")}
	ErrInvalidChar               = &Error{Op: "decode", Err: fmt.Errorf("invalid character found")}
	ErrUnknownConverter          = &Error{Op: "create", Err: fmt.Errorf("unknown converter type")}
	ErrConverterAlreadyRegistered = &Error{Op: "register", Err: fmt.Errorf("converter is already registered")}
	ErrConverterNotRegistered    = &Error{Op: "get", Err: fmt.Errorf("converter is not registered")}
)

func NewInvalidCharError(char rune) error {
	return &Error{Op: "decode", Err: fmt.Errorf("invalid character: %q", char)}
}

func NewUnknownConverterError(name string) error {
	return &Error{Op: "create", Err: fmt.Errorf("unknown converter type: %s", name)}
}
