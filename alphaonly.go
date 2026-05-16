package radixconverter

const alphaOnlyCharset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

type AlphaOnly struct {
	charset   string
	base      uint64
	charIndex map[rune]int
}

func NewAlphaOnly() *AlphaOnly {
	return &AlphaOnly{
		charset:   alphaOnlyCharset,
		base:      uint64(len(alphaOnlyCharset)),
		charIndex: BuildIndex(alphaOnlyCharset),
	}
}

func (a *AlphaOnly) Encode(number uint64) string {
	return Encode(number, a.charset, a.base)
}

func (a *AlphaOnly) Decode(encoded string) (uint64, error) {
	return Decode(encoded, a.base, a.charIndex)
}

func (a *AlphaOnly) Charset() string {
	return a.charset
}
