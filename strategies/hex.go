package strategies

const hexCharset = "0123456789abcdef"

// Hex implements a base-16 (hexadecimal) converter.
type Hex struct {
	charset   string
	base      uint64
	charIndex map[rune]int
}

// NewHex creates a new Hex converter.
func NewHex() *Hex {
	return &Hex{
		charset:   hexCharset,
		base:      uint64(len(hexCharset)),
		charIndex: BuildIndex(hexCharset),
	}
}

// Encode converts a non-negative integer to its hexadecimal string representation.
func (h *Hex) Encode(number uint64) string {
	return Encode(number, h.charset, h.base)
}

// Decode converts a hexadecimal string back to its integer value.
func (h *Hex) Decode(encoded string) (uint64, error) {
	return Decode(encoded, h.base, h.charIndex)
}

// Charset returns the charset used by this converter.
func (h *Hex) Charset() string {
	return h.charset
}
