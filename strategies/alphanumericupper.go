package strategies

const alphanumericUpperCharset = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// AlphanumericUpper implements a base-36 converter using digits and uppercase letters.
type AlphanumericUpper struct {
	charset   string
	base      uint64
	charIndex map[rune]int
}

// NewAlphanumericUpper creates a new AlphanumericUpper converter.
func NewAlphanumericUpper() *AlphanumericUpper {
	return &AlphanumericUpper{
		charset:   alphanumericUpperCharset,
		base:      uint64(len(alphanumericUpperCharset)),
		charIndex: BuildIndex(alphanumericUpperCharset),
	}
}

// Encode converts a non-negative integer to its alphanumeric-uppercase string representation.
func (a *AlphanumericUpper) Encode(number uint64) string {
	return Encode(number, a.charset, a.base)
}

// Decode converts an alphanumeric-uppercase string back to its integer value.
func (a *AlphanumericUpper) Decode(encoded string) (uint64, error) {
	return Decode(encoded, a.base, a.charIndex)
}

// Charset returns the charset used by this converter.
func (a *AlphanumericUpper) Charset() string {
	return a.charset
}
