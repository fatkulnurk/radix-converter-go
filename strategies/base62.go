package strategies

const base62Charset = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// Base62 implements a base-62 converter using digits, lowercase, and uppercase letters.
type Base62 struct {
	charset   string
	base      uint64
	charIndex map[rune]int
}

// NewBase62 creates a new Base62 converter.
func NewBase62() *Base62 {
	return &Base62{
		charset:   base62Charset,
		base:      uint64(len(base62Charset)),
		charIndex: BuildIndex(base62Charset),
	}
}

// Encode converts a non-negative integer to its base-62 string representation.
func (b *Base62) Encode(number uint64) string {
	return Encode(number, b.charset, b.base)
}

// Decode converts a base-62 string back to its integer value.
func (b *Base62) Decode(encoded string) (uint64, error) {
	return Decode(encoded, b.base, b.charIndex)
}

// Charset returns the charset used by this converter.
func (b *Base62) Charset() string {
	return b.charset
}
