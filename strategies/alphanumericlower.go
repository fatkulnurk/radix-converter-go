package strategies

const alphanumericLowerCharset = "0123456789abcdefghijklmnopqrstuvwxyz"

// AlphanumericLower implements a base-36 converter using digits and lowercase letters.
type AlphanumericLower struct {
	charset   string
	base      uint64
	charIndex map[rune]int
}

// NewAlphanumericLower creates a new AlphanumericLower converter.
func NewAlphanumericLower() *AlphanumericLower {
	return &AlphanumericLower{
		charset:   alphanumericLowerCharset,
		base:      uint64(len(alphanumericLowerCharset)),
		charIndex: BuildIndex(alphanumericLowerCharset),
	}
}

// Encode converts a non-negative integer to its alphanumeric-lowercase string representation.
func (a *AlphanumericLower) Encode(number uint64) string {
	return Encode(number, a.charset, a.base)
}

// Decode converts an alphanumeric-lowercase string back to its integer value.
func (a *AlphanumericLower) Decode(encoded string) (uint64, error) {
	return Decode(encoded, a.base, a.charIndex)
}

// Charset returns the charset used by this converter.
func (a *AlphanumericLower) Charset() string {
	return a.charset
}
