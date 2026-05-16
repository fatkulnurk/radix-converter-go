package radixconverter

const base62Charset = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

type Base62 struct {
	charset   string
	base      uint64
	charIndex map[rune]int
}

func NewBase62() *Base62 {
	return &Base62{
		charset:   base62Charset,
		base:      uint64(len(base62Charset)),
		charIndex: BuildIndex(base62Charset),
	}
}

func (b *Base62) Encode(number uint64) string {
	return Encode(number, b.charset, b.base)
}

func (b *Base62) Decode(encoded string) (uint64, error) {
	return Decode(encoded, b.base, b.charIndex)
}

func (b *Base62) Charset() string {
	return b.charset
}
