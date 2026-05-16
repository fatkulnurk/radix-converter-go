package strategies

const alphaOnlyCharset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// AlphaOnly implements a base-52 converter using only letters (no digits).
type AlphaOnly struct {
	charset   string
	base      uint64
	charIndex map[rune]int
}

// NewAlphaOnly creates a new AlphaOnly converter.
func NewAlphaOnly() *AlphaOnly {
	return &AlphaOnly{
		charset:   alphaOnlyCharset,
		base:      uint64(len(alphaOnlyCharset)),
		charIndex: BuildIndex(alphaOnlyCharset),
	}
}

// Encode converts a non-negative integer to its alpha-only string representation.
func (a *AlphaOnly) Encode(number uint64) string {
	return Encode(number, a.charset, a.base)
}

// Decode converts an alpha-only string back to its integer value.
func (a *AlphaOnly) Decode(encoded string) (uint64, error) {
	return Decode(encoded, a.base, a.charIndex)
}

// Charset returns the charset used by this converter.
func (a *AlphaOnly) Charset() string {
	return a.charset
}
