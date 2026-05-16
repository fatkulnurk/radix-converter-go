package radixconverter

const alphanumericLowerCharset = "0123456789abcdefghijklmnopqrstuvwxyz"

type AlphanumericLower struct {
	charset   string
	base      uint64
	charIndex map[rune]int
}

func NewAlphanumericLower() *AlphanumericLower {
	return &AlphanumericLower{
		charset:   alphanumericLowerCharset,
		base:      uint64(len(alphanumericLowerCharset)),
		charIndex: BuildIndex(alphanumericLowerCharset),
	}
}

func (a *AlphanumericLower) Encode(number uint64) string {
	return Encode(number, a.charset, a.base)
}

func (a *AlphanumericLower) Decode(encoded string) (uint64, error) {
	return Decode(encoded, a.base, a.charIndex)
}

func (a *AlphanumericLower) Charset() string {
	return a.charset
}
