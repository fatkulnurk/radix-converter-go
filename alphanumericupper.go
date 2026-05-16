package radixconverter

const alphanumericUpperCharset = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

type AlphanumericUpper struct {
	charset   string
	base      uint64
	charIndex map[rune]int
}

func NewAlphanumericUpper() *AlphanumericUpper {
	return &AlphanumericUpper{
		charset:   alphanumericUpperCharset,
		base:      uint64(len(alphanumericUpperCharset)),
		charIndex: BuildIndex(alphanumericUpperCharset),
	}
}

func (a *AlphanumericUpper) Encode(number uint64) string {
	return Encode(number, a.charset, a.base)
}

func (a *AlphanumericUpper) Decode(encoded string) (uint64, error) {
	return Decode(encoded, a.base, a.charIndex)
}

func (a *AlphanumericUpper) Charset() string {
	return a.charset
}
